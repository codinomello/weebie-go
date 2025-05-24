import { initializeApp } from "https://www.gstatic.com/firebasejs/11.8.1/firebase-app.js";
import { getAnalytics } from "https://www.gstatic.com/firebasejs/11.8.1/firebase-analytics.js";
import {
    getAuth,
    signInWithEmailAndPassword,
    createUserWithEmailAndPassword,
    GoogleAuthProvider,
    GithubAuthProvider,
    FacebookAuthProvider,
    signInWithPopup
} from "https://www.gstatic.com/firebasejs/11.8.1/firebase-auth.js";

// Configuração do Firebase
const firebaseConfig = {
    apiKey: "AIzaSyAI0Tc7GssKwWwtVdrz6OaK6KFACx58N5U",
    authDomain: "weebie-go.firebaseapp.com",
    projectId: "weebie-go",
    storageBucket: "weebie-go.appspot.com",
    messagingSenderId: "321509944065",
    appId: "1:321509944065:web:199a546b7055f461ec4900",
    measurementId: "G-S5CG0CLRVS"
};

// Inicializa o Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

// Referência para o auth
const auth = getAuth(app);

// Formulário de login
const signinForm = document.getElementById("email-signin-form");
if (signinForm) {
    signinForm.addEventListener("submit", async (e) => {
        e.preventDefault();

        // Obter valores do formulário
        const formData = new FormData(signinForm);
        const userData = Object.fromEntries(formData.entries());

        // Validação básica
        if (!userData.email || !userData.password) {
            showMessage("Por favor, preencha e-mail e senha", "error");
            return;
        }

        try {
            // Login com Firebase
            const userCredential = await signInWithEmailAndPassword(auth, userData.email, userData.password);

            // Obter token do Firebase
            const idToken = await userCredential.user.getIdToken();

            // Enviar para o backend
            const response = await fetch("/api/auth/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${idToken}`
                },
                body: JSON.stringify({
                    id_token: idToken,
                    uid: userCredential.user.uid
                })
            });

            if (!response.ok) {
                throw new Error(await response.text());
            }

            // Sucesso - redirecionar
            showMessage("Login realizado com sucesso!", "success");
            setTimeout(() => window.location.href = "/", 2000);
        } catch (error) {
            console.error("Erro no login:", error);
            showMessage(parseAuthError(error), "error");
        }
    });
}

// Formulário de cadastro
const signupForm = document.getElementById("email-signup-form");
if (signupForm) {
    signupForm.addEventListener("submit", async (e) => {
        e.preventDefault();

        // Obter valores do formulário
        const formData = new FormData(signupForm);
        const userData = Object.fromEntries(formData.entries());

        // Validação básica
        if (!userData.email || !userData.password) {
            showMessage("Por favor, preencha e-mail e senha", "error");
            return;
        }

        try {
            // Criar usuário no Firebase Auth
            const userCredential = await createUserWithEmailAndPassword(auth, userData.email, userData.password);

            // Obter token do Firebase
            const idToken = await userCredential.user.getIdToken();

            // Enviar dados para o backend
            const response = await fetch("/api/auth/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${idToken}`
                },
                body: JSON.stringify({
                    ...userData,
                    id_token: idToken,
                    uid: userCredential.user.uid
                })
            });

            if (!response.ok) {
                // Se o backend falhar, deleta o usuário do Firebase
                await userCredential.user.delete();
                throw new Error(await response.text());
            }

            // Sucesso - redirecionar
            showMessage("Cadastro realizado com sucesso!", "success");
            setTimeout(() => window.location.href = "/", 2000);
        } catch (error) {
            console.error("Erro no cadastro:", error);
            showMessage(parseAuthError(error), "error");
        }
    });
}

// Função para lidar com login social
async function handleProviderLogin(provider) {
    try {
        let authProvider;

        switch (provider) {
            case 'google':
                authProvider = new GoogleAuthProvider();
                break;
            case 'github':
                authProvider = new GithubAuthProvider();
                break;
            case 'facebook':
                authProvider = new FacebookAuthProvider();
                break;
            default:
                throw new Error("Provedor não suportado");
        }

        const result = await signInWithPopup(auth, authProvider);
        const user = result.user;
        const idToken = await user.getIdToken();

        // Enviar para seu backend
        const response = await fetch("/api/auth/social-login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${idToken}`
            },
            body: JSON.stringify({
                id_token: idToken,
                provider: provider,
                uid: user.uid,
                email: user.email,
                name: user.displayName
            })
        });

        if (!response.ok) {
            throw new Error(await response.text());
        }

        // Redirecionar após login bem-sucedido
        window.location.href = "/";
    } catch (error) {
        console.error("Erro no login social:", error);
        showMessage(parseAuthError(error), "error");
    }
}

// Função auxiliar para tratar erros de autenticação
function parseAuthError(error) {
    if (error.code) {
        switch (error.code) {
            case 'auth/email-already-in-use':
                return 'Este e-mail já está cadastrado.';
            case 'auth/invalid-email':
                return 'E-mail inválido.';
            case 'auth/weak-password':
                return 'Senha muito fraca (mínimo 6 caracteres).';
            case 'auth/popup-closed-by-user':
                return 'Login cancelado pelo usuário.';
            case 'auth/wrong-password':
                return 'Senha incorreta.';
            case 'auth/user-not-found':
                return 'Usuário não encontrado.';
            default:
                return `Erro: ${error.message}`;
        }
    }
    return error.message || 'Ocorreu um erro desconhecido.';
}

// Função para exibir mensagens
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