<template>
  <div v-bind:class="{'a': true}" class="b">
    <span class="d" v-bind:class="{c: true}" :a="1">
        {{data.msg}}
    </span>

    <div v-for="item in data.c">
      <bench :data="item"></bench>
    </div>
  </div>
</template>