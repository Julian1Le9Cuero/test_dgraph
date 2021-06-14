<template>
  <v-app id="inspire">
    <v-navigation-drawer 
    v-model="drawer"
    app>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="text-h6">
            Dgraph Shopper <v-icon color="grey lighten-1">mdi-cart-arrow-right</v-icon>
          </v-list-item-title>
          <v-list-item-subtitle>
            Shop fast
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
        dense
        nav
      >
        <v-list-item
          v-for="item in items"
          :key="item.title"
          :to="item.to"
          link
        >
          <v-list-item-icon>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-icon>

          <v-list-item-content>
            <v-list-item-title>{{ item.title }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

   <v-app-bar
      color="deep-purple accent-4"
      dense
      dark
      app
    >
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>Dgraph Shopper <v-icon color="white lighten-1">mdi-cart-arrow-right</v-icon></v-toolbar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-menu
        left
        bottom
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            icon
            v-bind="attrs"
            v-on="on"
          >
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="n in 5"
            :key="n"
            @click="() => {}"
          >
            <v-list-item-title>Option {{ n }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
    </v-app-bar>

    <v-main>
      <!-- Set routes -->
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
  export default {
    data: () => ({ 
      drawer: null,
      items: [
          { title: 'Dashboard', icon: 'mdi-view-dashboard', to: '/' },
          { title: 'Buyers', icon: 'mdi-cart-outline', to: '/buyers'},
        ]
     })
  }
</script>

<style>
p.error{
  border: 1px solid #ff5b5f;
  background-color: #ffc5c1;
  padding: 10px;
  margin-bottom: 15px;
}

.item-grid{
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.text-center{
  text-align: centers;
}

p.product-price{
  color: #39a9cb;
}

.trnsact-id{
  color: #e5d549;
}
</style>