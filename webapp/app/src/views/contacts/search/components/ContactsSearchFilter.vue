<template>
  <b-form @submit.prevent="onFilter">
    <b-form-row>
      <b-col>
        <b-form-group label="Nome">
          <b-input v-model="name" />
        </b-form-group>
      </b-col>

      <b-col>
        <b-form-group label="E-mail">
          <b-input
            type="email"
            v-model="$v.email.$model" />
          <b-form-invalid-feedback :force-show="$v.email.email === false">
            E-mail inválido
          </b-form-invalid-feedback>
        </b-form-group>
      </b-col>

      <b-col>
        <b-form-group label="Celular">
          <masked-input
            class="form-control"
            v-model="phone"
            :mask="phoneMask" />
        </b-form-group>
      </b-col>
    </b-form-row>
    <b-form-row>
      <b-col class="text-right">
        <b-button
          @click="onReset"
          variant="link"
          type="reset">
          Limpar
        </b-button>

        <b-button
          :disabled="$v.$invalid"
          variant="primary"
          type="submit">
          Pesquisar
        </b-button>
      </b-col>
    </b-form-row>
  </b-form>
</template>

<script>
  import email from 'vuelidate/lib/validators/email';
  import { PhoneMask } from '@/utils/masks';

  const LOCALSTORAGE_KEY = 'contacts/filter';

  export default {
    name: 'ContactsSearchFilter',
    data,
    validations,
    methods: {
      onFilter,
      onReset
    },
    mounted: onFilter
  };

  function data() {
    let savedFilter = localStorage.getItem(LOCALSTORAGE_KEY) || {};

    try {
      savedFilter = JSON.parse(savedFilter);
    } catch (e) {
      savedFilter = {};
    }

    return {
      name: savedFilter.name || null,
      email: savedFilter.email || null,
      phone: savedFilter.phone || null,
      phoneMask: PhoneMask
    };
  }

  function validations() {
    const validations = {
      email: {}
    };

    if (this.email) {
      validations.email.email = email;
    }

    return validations;
  }

  function onFilter() {
    const filter = {
      name: this.name,
      email: this.email,
      phone: this.phone
    }

    localStorage.setItem(LOCALSTORAGE_KEY, JSON.stringify(filter));

    this.$emit('filter', filter);
  }

  function onReset() {
    localStorage.removeItem(LOCALSTORAGE_KEY);
    this.name = null;
    this.email = null;
    this.phone = null;
    this.$emit('reset');
  }
</script>
