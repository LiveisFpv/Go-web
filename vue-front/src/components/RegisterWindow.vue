<script setup lang="ts">
  const register = () => {
    const username = (document.getElementById('username') as HTMLInputElement).value;
    const email = (document.getElementById('email') as HTMLInputElement).value;
    const password = (document.getElementById('password') as HTMLInputElement).value;
    console.log('Registering:', { username, email, password });
  };
  const goToLogin = () => {
    console.log('Redirecting to login page');  };
</script>

<template>
  <div id ="register-window">
    <h1>Register</h1>
    <form @submit.prevent="register">
      <div>
        <label for="username">Username:</label>
        <input type="text" id="email" placeholder="Email">
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="text" id="login" placeholder="Login">
      </div>
      <div>
        <label for="password">Password:</label>
       <input type="password" id="password" placeholder="Password">
      </div>
      <button type="submit">Register</button>
      <button type="button" @click="goToLogin">Already have an account? Login</button>
      <p id="text">By registering, you agree to our <a href="#">Terms of Service</a> and <a href="#">Privacy Policy</a>.</p>
    </form>
  </div>
</template>
<style scoped>

</style>
