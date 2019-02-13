import Vue from 'vue';
import DefaultLayout from './../layouts/DefaultLayout.vue';
import TitleLayout from '../layouts/TitleLayout.vue';
import MaskedInput from 'vue-text-mask';

Vue.component(DefaultLayout.name, DefaultLayout);
Vue.component(TitleLayout.name, TitleLayout);
Vue.component(MaskedInput.name, MaskedInput);