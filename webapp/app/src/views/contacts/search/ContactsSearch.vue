<template>
  <TitleLayout title="Agenda de Contatos">
    <b-row>
      <b-col>
        <ContactsSearchFilter
          @filter="onFilter"
          @reset="onFilter(null)" />
      </b-col>
    </b-row>

    <hr />

    <b-row class="mb-3">
      <b-col>
        Total: {{ totalCount }}
      </b-col>
      <b-col class="text-right">
        <b-button
          size="sm"
          variant="outline-secondary"
          :to="{ name: 'contacts-new' }">
          Cadastrar novo
        </b-button>
      </b-col>
    </b-row>

    <b-row>
      <b-col>
        <ContactsSearchTable
          :sorting="sort"
          :contacts="contacts"
          @sort="onSort" />

        <b-pagination
          v-if="totalPages > 1"
          :total-rows="totalCount"
          :per-page="limit" />
      </b-col>
    </b-row>
  </TitleLayout>
</template>

<script>
  import axios from 'axios';
  import ContactsSearchFilter from './components/ContactsSearchFilter.vue';
  import ContactsSearchTable from './components/ContactsSearchTable.vue';

  export default {
    name: 'ContactsSearch',
    components: { ContactsSearchFilter, ContactsSearchTable },
    data,
    methods: {
      onFilter,
      toPage,
      onSort,
      _getContactPage
    }
  };

  function data() {
    return {
      filter: null,
      contacts: [],
      page: 0,
      sort: null,
      limit: 10,
      totalCount: 0,
      totalPages: 0
    };
  }

  function onFilter(filter) {
    this.page = 0;
    this.sort = null;
    this.filter = filter;
    this._getContactPage();
  }

  function toPage(page) {
    this.page = page;
    this._getContactPage();
  }

  function onSort(sort) {
    this.sort = sort;
    this._getContactPage();
  }

  function _getContactPage() {
    const params = {
      _page: this.page,
      _limit: this.limit
    };

    if (this.sort) {
      params._sort = this.sort.by;
      params._order = this.sort.order;
    }

    if (this.filter != null) {
      Object.keys(this.filter).forEach(key => params[key] = this.filter[key]);
    }

    axios.get('/contacts', { params })
      .then(({ data, headers }) => {
        this.totalCount = headers['x-total-count'];
        this.totalPages = this.totalCount / this.limit;
        this.contacts = data;
      });
  }
</script>
