<template>
    <span>
        <v-app-bar >
            <v-btn style="margin:10px;" to="/home">{{ appTitle }}</v-btn>
            <h4>{{ $t('headerText') }}</h4>
            <v-spacer class="hidden-md-and-up"></v-spacer>
            <v-spacer class="hidden-sm-and-down"></v-spacer>

            <div class="locale-changer">
                <v-select
                        v-model="$i18n.locale"
                        :items="locales"
                        menu-props="auto"
                        label="Select"
                        hide-details
                        prepend-icon="map"
                        single-line
                ></v-select>
             </div>
        </v-app-bar>
    </span>
</template>

<script>

    export default {
        name: "AppNavigation",
        data() {
            return {
                appTitle: 'Warte Warte',
                drawer: false,
                items: [
                    {title: 'Home'},
                ],
                locales: this.$i18n.availableLocales
            };
        },
        created() {
            let browserLocale = this.getBrowserLocale(this.$i18n.availableLocales, this.$i18n.fallbackLocale);
            console.info("setting locale to: " + browserLocale);
            this.$i18n.locale = browserLocale;

           /*
           todo this should be unit tests!
           console.info("de01: " + this.matchBrowserLocale(['de', 'en'], 'en', ['de_DE']));
           console.info("de02: " + this.matchBrowserLocale(['de', 'en'], 'en', ['de']));
           console.info("de03: " + this.matchBrowserLocale(['de', 'en'], 'en', ['xx_YY', 'de']));

           console.info("en01: " + this.matchBrowserLocale(['de', 'en'], 'en', ['xx_YY']));
           console.info("en02: " + this.matchBrowserLocale(['de', 'en'], 'en', ['en_US']));
           */
        },
        methods: {
            getBrowserLocale: function (possibleLocales, defaultLocale) {

                let browserLanguages;
                if (navigator.languages !== undefined) {
                    browserLanguages = navigator.languages;
                } else {
                    browserLanguages = [navigator.language];
                }

                return this.matchBrowserLocale(possibleLocales, defaultLocale, browserLanguages);
            },
            matchBrowserLocale: function (possibleLocales, defaultLocale, browserLanguages) {

                for (let i = 0; i < browserLanguages.length; i++) {
                    let trimmed = browserLanguages[i].trim().split(/[-_]/)[0];
                    if (possibleLocales.includes(trimmed)) {
                        return trimmed;
                    }
                }
                return defaultLocale;
            }
        }
    }
</script>

<style scoped>

    .locale-changer {
        width: 100px;
    }

</style>