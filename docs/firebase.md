# Firebase 🔐

## Introdução
Firebase Authentication é um serviço que facilita a implementação de sistemas de autenticação, suportando diversos provedores como e-mail/senha, Google, Facebook, entre outros.

## Configuração
- Acesse o [Firebase Console](https://firebase.google.com/products/auth) e crie um projeto.
- Ative os provedores de autenticação desejados e configure as regras de segurança.

## Conceitos básicos
- **Integração Rápida:** Permite a implementação rápida de autenticação sem complicações.
- **Segurança:** Garante o gerenciamento seguro de usuários e credenciais.
- **Escalabilidade:** Adequado para projetos de qualquer tamanho, do iniciante ao enterprise.

## Exemplo de uso
```javascript
// Exemplo básico de autenticação utilizando Firebase Auth com JavaScript
import { initializeApp } from "firebase/app";
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";

const firebaseConfig = {
  apiKey: "SUA_API_KEY",
  authDomain: "SEU_PROJETO.firebaseapp.com",
  // Outras configurações...
};

const app = initializeApp(firebaseConfig);
const auth = getAuth(app);

signInWithEmailAndPassword(auth, "usuario@exemplo.com", "senha123"){
  .then((userCredential) => {
    const user = userCredential.user;
    console.log("Usuário autenticado:", user);
  })
  .catch((error) => {
    console.error("Erro na autenticação:", error);
  });
};