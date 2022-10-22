<template>
    <b-card :title="profile.name+' (edit mode)'">
        <b-form-group label="Fronting alters:">
            <b-form-checkbox
                v-for="(alter, index) in profile.alters"
                :key="alter.title"
                :checked="alter.fronting"
                @change="setAlter(index); save()"
                >
                {{ alter.title }}
            </b-form-checkbox>
        </b-form-group>
        <b-button variant="outline-primary">Add alter</b-button>
        <hr/>

        <b-form-group
            id="input-group-tz"
            label="Timezone:"
            label-for="input-tz"
            >
            <b-form-input
                id="input-tz"
                v-model="profile.timezone"
                type="text"
                required
                @change="save()"
            ></b-form-input>
        </b-form-group>
        <hr/>

        <b-form-group label="Warnings/Notices">
            <b-form-textarea
                id="notices"
                v-model="notices"
                placeholder="One notice per line"
                rows="3"
                max-rows="6"
                @change="save()"
                ></b-form-textarea>
        </b-form-group>

        <hr/>

        <LightEdit v-for="(light, index) in profile.lights" v-model="profile.lights[index]"   @change="save()"/>
        <b-button variant="outline-primary">Add indicator light</b-button>
    </b-card>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'EditProfile',
    data() {
        return {
            profile: {
                alters: [],
                lights: [],
                notices: [],
            },
            now: Date.now(),
            date_formatter: new Intl.DateTimeFormat(),
        }
    },
    computed: {
        current_time() {
            return this.date_formatter.format(this.now)
        },
        fronting_alters() {
            return this.profile.alters.filter(alter => alter.fronting).map(alter => alter.title)
        },
        notices: {
            get() {
                return this.profile.notices.join('\n\n')
            },
            set(new_value) {
                this.profile.notices = new_value.split('\n\n')
            }
        }
    },
    created() {
        var self = this
        setInterval(function () {
            self.now = Date.now()
        }, 1000)
    },
    async fetch() {
        const profile_ulid = this.$route.params.ulid;
        this.profile = await fetch('http://localhost:3001/profiles/'+profile_ulid).then(res => res.json())
        console.log(this.profile)
        this.date_formatter = new Intl.DateTimeFormat([], {
            timeZone: this.profile.timezone,
            year: 'numeric',
            month: 'numeric',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric',
        });
    },
    methods: {
        setAlter(index) {
            console.log("toggle", index)
            this.profile.alters[index].fronting = !this.profile.alters[index].fronting
            console.log("toggle", this.profile.alters)
        },
        save() {
            this.$axios.put('http://localhost:3001/profiles/'+this.profile.ulid, this.profile).then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });
        }
    }
})
</script>