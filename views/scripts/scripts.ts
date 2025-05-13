// URL do backend
const backendUrl = "http://localhost:8080";

// Função para cadastrar usuário
const registerForm = document.getElementById("registerForm") as HTMLFormElement;
const emailInput = document.getElementById("email") as HTMLInputElement;
const passwordInput = document.getElementById("password") as HTMLInputElement;
const nameInput = document.getElementById("name") as HTMLInputElement;
const registerMessage = document.getElementById("registerMessage") as HTMLParagraphElement;

registerForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const email = emailInput.value;
    const password = passwordInput.value;
    const name = nameInput.value;

    try {
        const response = await fetch(`${backendUrl}/signup`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password, name }),
        });

        const result = await response.json();

        if (response.ok) {
            registerMessage.textContent = `Usuário cadastrado com sucesso! UID: ${result.uid}`;
            registerMessage.style.color = "green";
        } else {
            registerMessage.textContent = `Erro: ${result.error || "Falha no cadastro"}`;
            registerMessage.style.color = "red";
        }
    } catch (error) {
        registerMessage.textContent = "Erro ao conectar ao servidor";
        registerMessage.style.color = "red";
    }
});

// Função para cadastrar projeto
const projectForm = document.getElementById("project-form") as HTMLFormElement;
const ownerUidInput = document.getElementById("owner-uid") as HTMLInputElement;
const titleInput = document.getElementById("title") as HTMLInputElement;
const detailsInput = document.getElementById("details") as HTMLTextAreaElement;
const groupInput = document.getElementById("group") as HTMLInputElement;
const yearInput = document.getElementById("year") as HTMLInputElement;
const courseInput = document.getElementById("course") as HTMLInputElement;
const odsInput = document.getElementById("ods") as HTMLInputElement;
const projectMessage = document.getElementById("project-message") as HTMLParagraphElement;

projectForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const project = {
        owner_uid: ownerUidInput.value,
        title: titleInput.value,
        details: detailsInput.value,
        group: groupInput.value,
        year: parseInt(yearInput.value),
        course: courseInput.value,
        ods: odsInput.value.split(",").map(Number),
    };

    try {
        const response = await fetch(`${backendUrl}/project`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(project),
        });

        const result = await response.json();

        if (response.ok) {
            projectMessage.textContent = `Projeto cadastrado com sucesso! ID: ${result.id}`;
            projectMessage.style.color = "green";
        } else {
            projectMessage.textContent = `Erro: ${result.error || "Falha no cadastro"}`;
            projectMessage.style.color = "red";
        }
    } catch (error) {
        projectMessage.textContent = "Erro ao conectar ao servidor";
        projectMessage.style.color = "red";
    }
});