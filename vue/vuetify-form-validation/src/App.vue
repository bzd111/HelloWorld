<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <v-toolbar-title>Vuetify Form Validation</v-toolbar-title>
    </v-app-bar>

    <v-content>
      <v-tabs v-model="tab" centered>
        <v-tab>Login</v-tab>
        <v-tab>Registration</v-tab>
      </v-tabs>
      <v-tabs-items v-model="tab">
       <v-tab-item name="test1">
      <LoginForm></LoginForm>
      </v-tab-item> 
      <v-tab-item><RegistrationForm></RegistrationForm></v-tab-item>
      </v-tabs-items> 
    </v-content>
  </v-app>
</template>

<script>
import LoginForm from "./components/LoginForm";
import RegistrationForm from "./components/RegistrationForm"

export default {
  name: "App",

  components: {
    LoginForm,
    RegistrationForm
  },

  data: () => ({
    tab: null
  })
};
</script>
