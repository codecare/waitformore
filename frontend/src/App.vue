<template>
    <v-app>
        <app-navigation></app-navigation>

        <v-content transition="slide-x-transition">
            <router-view></router-view>
        </v-content>

        <legal-requirements-footer></legal-requirements-footer>
    </v-app>
</template>

<script>
    import AppNavigation from "./components/AppNavigation";
    import LegalRequirementsFooter from "./components/LegalRequirementsFooter";

    export default {
        name: 'App',
        components: {
            AppNavigation,
            LegalRequirementsFooter
        }
    }
</script>