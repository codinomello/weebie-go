  // Import the functions you need from the SDKs you need
  import { initializeApp } from "https://www.gstatic.com/firebasejs/11.3.1/firebase-app.js";
  import { getAnalytics } from "https://www.gstatic.com/firebasejs/11.3.1/firebase-analytics.js";
  // TODO: Add SDKs for Firebase products that you want to use
  // https://firebase.google.com/docs/web/setup#available-libraries

  // Your web app's Firebase configuration
  // For Firebase JS SDK v7.20.0 and later, measurementId is optional
  const firebaseConfig = {
    apiKey: "AIzaSyAI0Tc7GssKwWwtVdrz6OaK6KFACx58N5U",
    authDomain: "weebie-go.firebaseapp.com",
    projectId: "weebie-go",
    storageBucket: "weebie-go.firebasestorage.app",
    messagingSenderId: "321509944065",
    appId: "1:321509944065:web:199a546b7055f461ec4900",
    measurementId: "G-S5CG0CLRVS"
  };

  // Initialize Firebase
  const app = initializeApp(firebaseConfig);
  const analytics = getAnalytics(app);
