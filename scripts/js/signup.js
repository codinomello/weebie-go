import { getAuth, createUserWithEmailAndPassword } from 'https://www.gstatic.com/firebasejs/11.8.0/firebase-auth.js';
            
const auth = getAuth();
            const signupForm = document.getElementById("email-signup-form");
            
            signupForm?.addEventListener("submit", async (e) => {
                e.preventDefault();
                
                // Get form data
                const formData = new FormData(signupForm);
                const userData = Object.fromEntries(formData.entries());
                
                // Basic validation
                if (!userData.email || !userData.password) {
                    showMessage("Por favor, preencha e-mail e senha", "error");
                    return;
                }

                try {
                    // Create user with Firebase Auth
                    const userCredential = await createUserWithEmailAndPassword(
                        auth,
                        userData.email,
                        userData.password
                    );
                    
                    // Get Firebase ID token
                    const idToken = await userCredential.user.getIdToken();
                    
                    // Prepare complete user data
                    const completeUserData = {
                        ...userData,
                        id_token: idToken,
                        uid: userCredential.user.uid,
                        role: "user",
                        status: "active"
                    };

                    // Send to backend
                    const response = await fetch("/api/auth/register", {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json",
                            "Authorization": `Bearer ${idToken}`
                        },
                        body: JSON.stringify(completeUserData)
                    });

                    if (!response.ok) {
                        // If backend fails, delete the Firebase user
                        await userCredential.user.delete();
                        throw new Error(await response.text());
                    }

                    // Success - redirect
                    showMessage("Cadastro realizado com sucesso!", "success");
                    setTimeout(() => window.location.href = "/", 2000);
                } catch (error) {
                    console.error("Erro no cadastro:", error);
                    showMessage(parseAuthError(error), "error");
                }
            });

            // Helper function to parse auth errors
            function parseAuthError(error) {
                if (error.code) {
                    switch (error.code) {
                        case 'auth/email-already-in-use':
                            return 'Este e-mail já está cadastrado.';
                        case 'auth/invalid-email':
                            return 'E-mail inválido.';
                        case 'auth/weak-password':
                            return 'Senha muito fraca (mínimo 6 caracteres).';
                        case 'auth/operation-not-allowed':
                            return 'Operação não permitida.';
                        default:
                            return `Erro: ${error.message}`;
                    }
                }
                return error.message || 'Ocorreu um erro desconhecido.';
            }

            // Function to show messages
            function showMessage(message, type) {
                const authResult = document.getElementById("auth-result");
                if (!authResult) return;

                authResult.innerHTML = `
                    <div class="p-3 mb-4 text-sm rounded-lg ${
                        type === 'success' 
                            ? 'bg-green-100 text-green-700 dark:bg-green-200 dark:text-green-800' 
                            : 'bg-red-100 text-red-700 dark:bg-red-200 dark:text-red-800'
                    }">
                        ${message}
                    </div>
                `;
                authResult.classList.remove("hidden");
            }