<template>
<div class="form-container">
    <h2>Login</h2>
    <form @submit.prevent="Login">
        <div class="form-group">
            <label for="email">Email</label>
            <input type="email" v-model="FormData.email" id="email" name="email" required>
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input type="password" v-model="FormData.password" id="password" name="password" required >
        </div>
        <button type="submit">Login</button>
    </form>
</div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'LOGIN',
    data() {
        return {
            FormData: {
                email: '',
                password: ''
            },
            message: ''
        }
       
    },
    methods: {
        async Login() {
            try {
                const response = await axios.post('http://localhost:8000/api/login', this.FormData);
                this.message = response.data.message;
                alert("User registered successfully!");
            } catch (error) {
                this.message = 'An error occurred while submitting the form.';
                alert(this.message);
            }

        }
    }

}
</script>

<style>
body {
    font-family: Arial, sans-serif;
    background-color: #f4f4f4;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    margin: 0;
}

.form-container {
    background: #fff;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    width: 300px;
    text-align: center;
}

.form-group {
    margin-bottom: 15px;
    text-align: left;
}

label {
    display: block;
    font-weight: bold;
    margin-bottom: 5px;
}

input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
    font-size: 16px;
    box-sizing: border-box;
}

button {
    width: 100%;
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    font-size: 16px;
    cursor: pointer;
    transition: 0.3s;
}

button:hover {
    background-color: #0056b3;
}
</style>
