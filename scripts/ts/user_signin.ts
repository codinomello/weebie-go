const form = document.getElementById("signup-form");

if (form) {
  	form.addEventListener("submit", async function (e) {
    	e.preventDefault();

		const payload = {
			name: (document.getElementById("name") as HTMLInputElement).value,
			email: (document.getElementById("email") as HTMLInputElement).value,
			password: (document.getElementById("password") as HTMLInputElement).value,
		};

		try {
			const response = await fetch("http://localhost:8080/users", {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify(payload)
			});
			const result = await response.json();
			console.log("Usuário criado:", result);
			// Mensagem de sucesso aqui :)
		} catch (err) {
			console.error("Erro ao cadastrar:", err);
			// Mensagem de erro aqui :(
		}
 	 });
}