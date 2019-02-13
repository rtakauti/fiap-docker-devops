import Vue from 'vue';
import DefaultLayout from './../layouts/DefaultLayout.vue';
import TituloLayout from './../layouts/TituloLayout.vue';
import MaskedInput from 'vue-text-mask';

Vue.component(DefaultLayout.name, DefaultLayout);
Vue.component(TituloLayout.name, TituloLayout);
Vue.component(MaskedInput.name, MaskedInput);