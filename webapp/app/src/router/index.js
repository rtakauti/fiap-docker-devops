import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

export const Router = new VueRouter({
  mode: 'history',
  routes: [
    {
      name: 'contacts',
      path: '/contacts',
      component: () => import(/* webpackChunkName: "ContactsSearch" */ '@/views/contacts/search/ContactsSearch.vue')
    },
    {
      name: 'contacts-new',
      path: '/contacts/new',
      component: () => import(/* webpackChunkName: "ContactForm" */ '@/views/contacts/form/ContactForm.vue'),
      meta: {
        title: 'Novo Contato'
      }
    },
    {
      name: 'contacts-edit',
      path: '/contacts/:id',
      component: () => import(/* webpackChunkName: "ContactForm" */ '@/views/contacts/form/ContactForm.vue'),
      props: true,
      meta: {
        title: 'Editar Contato'
      }
    },
    {
      path: '**',
      redirect: {
        name: 'contacts'
      }
    }
  ]
});