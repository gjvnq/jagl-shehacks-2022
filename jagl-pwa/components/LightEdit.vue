<template>
    <div class="light-edit">
        <b-form-group label="Indicator label">
            <b-form-input v-model="value.title" @input="oninput()" @change="onchange()"  placeholder="Indicator title"></b-form-input>
        </b-form-group>
        <b-form-group label="Current status" v-slot="{ ariaDescribedby }">
            <b-form-radio v-model.number="value.color" @input="oninput()" @change="onchange()" :aria-describedby="ariaDescribedby" :name="'color-status-'+rand_id" value="0" class="light-0">Muted</b-form-radio>
            <b-form-radio v-model.number="value.color" @input="oninput()" @change="onchange()" :aria-describedby="ariaDescribedby" :name="'color-status-'+rand_id" value="1" class="light-1">Improving</b-form-radio>
            <b-form-radio v-model.number="value.color" @input="oninput()" @change="onchange()" :aria-describedby="ariaDescribedby" :name="'color-status-'+rand_id" value="2" class="light-2">Good</b-form-radio>
            <b-form-radio v-model.number="value.color" @input="oninput()" @change="onchange()" :aria-describedby="ariaDescribedby" :name="'color-status-'+rand_id" value="3" class="light-3">Danger</b-form-radio>
            <b-form-radio v-model.number="value.color" @input="oninput()" @change="onchange()" :aria-describedby="ariaDescribedby" :name="'color-status-'+rand_id" value="4" class="light-4">Crisis</b-form-radio>
        </b-form-group>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'LightEdit',
    props: ['value'],
    data() {
        return {
            rand_id: Math.round(Math.random()*1000)
        }
    },
    methods: {
        onchange() {
            this.$emit('change', this.value)
        },
        oninput() {
            this.$emit('input', this.value)
        }
    }
})
</script>

<style>
.light-edit {
    display: inline-block;
    margin: 0.1rem 0.5rem;
}
.light-0 {
    color: gray;
}
.light-1 {
    color: blue;
}
.light-2 {
    color: green;
}
.light-3 {
    color: orange;
}
.light-4 {
    color: red;
}
</style>