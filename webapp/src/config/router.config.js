import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

export const Router = new VueRouter({
  mode: 'history',
  routes: [
    {
      name: 'contatos',
      path: '/contatos',
      component: () => import(/* webpackChunkName: "ConsultaContatos" */ '@/views/contatos/consulta/ConsultaContatos.vue')
    },
    {
      name: 'contato-novo',
      path: '/contatos/novo',
      component: () => import(/* webpackChunkName: "NovoContato" */ '@/views/contatos/formulario/NovoContato.vue')
    },
    {
      name: 'contato-editar',
      path: '/contatos/:id',
      component: () => import(/* webpackChunkName: "EditarContato" */ '@/views/contatos/formulario/EditarContato.vue')
    },
    {
      path: '**',
      redirect: 'contatos'
    }
  ]
});