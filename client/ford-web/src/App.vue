<template>
  <nav>
    <router-link to="/">create new</router-link>
    <router-link to="/12.02.2022">12.02.2022</router-link>
  </nav>
  <router-view />
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: sans-serif;
  width: 100%;
  min-height: 100vh;
  background-color: #eeeff3;
}
</style>
