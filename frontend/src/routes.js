import Vue from 'vue'
import VueRouter from "vue-router";
import Home from "./components/Home";
import CreateQueue from "./components/CreateQueue";
import QueueAdmin from "./components/QueueAdmin";
import UserView from "./components/UserView";


Vue.use(VueRouter);

const routes = [
    { path: '/home', component: Home },
    { path: '/usage', component: Home },
    { path: '/createQueue', component: CreateQueue },
    { path: '/queueAdmin/:adminKey', component: QueueAdmin },
    { path: '/user/:userKey', component: UserView },
    { path: '/', redirect: '/home' }
];

export default new VueRouter({
    routes // short for `routes: routes`
});