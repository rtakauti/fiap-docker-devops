<template>
  <TitleLayout :title="title">
    <b-form-row>
      <b-col>
        <b-form @submit.prevent="onSubmit">
          <b-form-group label="Nome">
            <b-input v-model="$v.name.$model" />
            <b-form-invalid-feedback :force-show="$v.name.$dirty && $v.name.required === false">
              Nome é um campo obrigatório
            </b-form-invalid-feedback>
            <b-form-invalid-feedback :force-show="$v.name.$dirty && $v.name.minLength === false">
              Nome deve conter pelo menos 5 caractéres
            </b-form-invalid-feedback>
          </b-form-group>

          <b-form-group label="E-mail">
            <b-input
              type="email"
              v-model="$v.email.$model" />
            <b-form-invalid-feedback :force-show="$v.email.$dirty && $v.email.required === false">
              E-mail é um campo obrigatório
            </b-form-invalid-feedback>
            <b-form-invalid-feedback :force-show="$v.email.$dirty && $v.email.email === false">
              E-mail inválido
            </b-form-invalid-feedback>
          </b-form-group>

          <b-form-group label="Celular">
            <masked-input
              :mask="phoneMask"
              class="form-control"
              v-model="$v.phone.$model" />
            <b-form-invalid-feedback :force-show="$v.phone.$dirty && $v.phone.required === false">
              Celular é um campo obrigatório
            </b-form-invalid-feedback>
            <b-form-invalid-feedback :force-show="$v.phone.$dirty && $v.phone.minLength === false">
              Celular inválido
            </b-form-invalid-feedback>
          </b-form-group>

          <b-row>
            <b-col>
              <b-button
                v-if="isEditing"
                variant="danger"
                @click="onDelete">
                Deletar
              </b-button>
            </b-col>

            <b-col class="text-right">
              <b-button
                class="m-l-1 m-r-1"
                variant="link"
                :to="{ name: 'contacts' }">
                Cancelar
              </b-button>

              <b-button
                type="submit"
                variant="primary">
                Salvar
              </b-button>
            </b-col>
          </b-row>
        </b-form>
      </b-col>
    </b-form-row>
  </TitleLayout>
</template>

<script>
  import axios from 'axios';
  import Noty from 'noty';
  import { required, email, minLength } from 'vuelidate/lib/validators';
  import { PhoneMask } from '@/utils/masks';

  export default {
    name: 'ContactForm',
    props: ['id'],
    data,
    validations,
    computed: {
      title,
      isEditing
    },
    methods: {
      onDelete,
      onSubmit
    },
    mounted
  };

  function data() {
    return {
      contact: null,
      name: null,
      email: null,
      phone: null,
      phoneMask: PhoneMask
    };
  }

  function validations() {
    return {
      name: {
        required,
        minLength: minLength(5)
      },
      email: {
        required,
        email
      },
      phone: {
        required,
        phone: minLength(10)
      }
    };
  }

  function title() {
    if (this.isEditing && this.contact !== null) {
      return `${this.$route.meta.title} - ${this.contact.name}`;
    } else {
      return this.$route.meta.title;
    }
  }

  function isEditing() {
    return this.$route.name === 'contacts-edit';
  }

  function onDelete() {
    axios.delete(`/contacts/${this.id}`)
      .then(({ data }) => {
        new Noty({
          text: 'Contato deletado com sucesso',
          type: 'success'
        }).show();
        this.$router.replace({
          name: 'contacts'
        });
      });
  }

  function onSubmit() {
    this.$v.$touch();

    if (this.$v.$invalid) {
      new Noty({
        text: 'Preencha o formulário corretamente para poder salvar',
        type: 'warning'
      }).show();
    } else {
      axios.request({
        url: this.isEditing ? `/contacts/${this.id}` : '/contacts',
        method: this.isEditing ? 'PUT' : 'POST',
        data: {
          name: this.name,
          email: this.email,
          phone: this.phone
        }
      }).then(({ data }) => {
        if (this.isEditing) {
          new Noty({
            text: 'Contato atualizado com sucesso',
            type: 'success'
          }).show();

          this.$router.push({
            name: 'contacts'
          });
        } else {
          new Noty({
            text: 'Contato criado com sucesso',
            type: 'success'
          }).show();

          this.$router.replace({
            name: 'contacts-edit',
            params: {
              id: data.id
            }
          });
        }
      });
    }
  }

  function mounted() {
    if (this.isEditing) {
      axios.get(`/contacts/${this.id}`)
        .then(({ data }) => {
          this.contact = data;
          this.name = this.contact.name;
          this.email = this.contact.email;
          this.phone = this.contact.phone;
        })
        .catch(() => {
          new Noty({
            text: 'Contato não encontrado',
            type: 'error'
          }).show();
          this.$router.replace({
            name: 'contacts'
          });
        });
    }
  }
</script>
