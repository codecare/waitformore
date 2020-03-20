<template>
    <v-container>

        <v-row>
            <v-col>
                <h3>{{ $t('userQueueName') }}<br>
                    <span class="queueTitle"> {{ queueTitle }}</span></h3>
                <div v-if="$i18n.locale === 'de'">
                    <p>Jeder der möchte, kann diesen Service kostenlos verwenden <br>
                        Es ist jedoch eine Beta-Version - wir übernehmen keine Haftung oder Verantwortung bei Fehlern oder Ausfällen!</p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <p>This service is freely available to everyone. <br>
                        This is a Beta-version - we are very invested in making sure it all works, however we cannot be held responsible for failures</p>
                </div>
            </v-col>
        </v-row>

        <v-row>
            <v-col>
                <v-alert type="warning"
                         dismissible
                         transition="scale-transition"
                         :value="showTechnicalError">
                    {{ $t('technicalProblem') }}
                </v-alert>
                <v-alert color="blue" dark>
                    <v-row align="center">
                        <v-col class="shrink">
                            <img src="../assets/queueAdmin-white.png" :alt="$t('altQueueUp')" width="150" height="150"/>
                        </v-col>
                        <v-col class="grow"><span class="lastCalled">{{ queueState.LastCalled }}</span> {{ $t('userLastCalled') }}</v-col>
                    </v-row>
                </v-alert>
            </v-col>
        </v-row>

        <v-row>
            <v-col>

                <v-alert color="orange"
                         transition="scale-transition"
                         :value="showOwnNumberWasCalled"
                         class="userIsCalled">
                    <span class="showOwnNumberWasCalled">{{ $t('userIsCalled') }}!</span>
                </v-alert>

            </v-col>
        </v-row>

        <v-row>
            <v-col>
                <v-alert v-if="!userCanPull" color="green" dark>
                    <v-row align="center">
                        <v-col class="shrink">
                            <img src="../assets/user-white.png" alt="an Schlange anstellen" width="150" height="150"/>
                        </v-col>
                        <v-col class="grow"><span class="userPulledNumber">{{ userPulledNumber }}</span> {{ $t('userWaitingNumber') }}</v-col>
                    </v-row>
                </v-alert>

                <v-alert v-if="userCanPull" color="green" dark>
                    <v-row align="center">
                        <v-col class="shrink">
                            <img src="../assets/user-white.png" alt="an Schlange anstellen" width="150" height="150"/>
                        </v-col>
                        <v-col class="grow"><v-btn :disabled="!userCanPull" @click="pullNext">{{ $t('pullWaitingNumber' )}}</v-btn></v-col>
                    </v-row>
                </v-alert>
            </v-col>
        </v-row>

        <v-row>
            <v-col>
                <div v-if="$i18n.locale === 'de'">
                    <h3>Url zum Anstellen an der Warteschlange: {{ queueTitle }}</h3>
                    <p>Diesen können Sie an wartende Kunden weitergeben, sobald Sie eine Nummer gezogen haben.</p>
                    <p><a class="userKey" :href="userUrl">{{ userUrl }}</a></p>
                    <p>Alternativ kann der Code <span class="userKey">{{ userKey }}</span> auf der Homepage <a :href="baseUrl" class="baseUrl">{{ baseUrl }}</a> eingegeben werden.</p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <h3>Link to queue: {{ queueTitle }}</h3>
                    <p>You may pass this on to others after you have joined the queue - you will keep your place.</p>
                    <p><a class="userKey" :href="userUrl">{{ userUrl }}</a></p>
                    <p>You may also tell them to enter the code <span class="userKey">{{ userKey }}</span> on the home page <a c class="baseUrl">{{ baseUrl }}</a></p>
                </div>
            </v-col>
            <v-col>
                <a class="userKey" :href="userUrl" :target="userKey">
                    <qrcode-vue :value="userUrl" size="300" level="H"></qrcode-vue>
                    {{ userUrl }}
                </a>
            </v-col>
        </v-row>

        <v-row>
            <v-col>
                <v-btn @click="refresh">{{ $t('btnRefresh') }}</v-btn>
                <p>{{ $t('autoRefresh') }}</p>
            </v-col>
        </v-row>

    </v-container>
</template>

<script>

    import Vue from 'vue'
    import VueCookies from 'vue-cookies'
    import QrcodeVue from 'qrcode.vue'

    Vue.use(VueCookies)

    export default {
        name: "QueueAdmin",
        data: () => ({
            queueTitle: '',
            userKey: '',
            queueState: {},
            userPulledNumber: 0,
            userCanPull: true,
            timer: '',
            userUrl: '',
            baseUrl: '',
            showOwnNumberWasCalled: false,
            showTechnicalError: false
        }),
        components: {
            VueCookies,
            QrcodeVue,
        },
        created() {
            this.userKey = this.$route.params.userKey;

            let wasPulled = this.$cookies.get(this.userKey);

            if (wasPulled)
            {
                this.userPulledNumber = wasPulled;
                this.userCanPull = false;
            }
            this.userUrl = this.createUserUrl();
            this.baseUrl = this.createBaseUrl();
            this.refresh();
            this.timer = setInterval(this.refresh, 10000);
        },
        methods: {
            refresh: function() {
                console.info("refresh")
                fetch(process.env.VUE_APP_URL + '/user/' + this.userKey)
                    .then(response => {
                        if (response.status === 200) {
                            return response.json();
                        } else {
                            console.warn('response failed');
                            return {};
                        }
                    })
                    .catch(() => {
                        (console.warn("Network connection error"));
                        return {};
                    })
                    .then(json => {
                        if (json.QueueTitle) {
                            this.queueTitle = json.QueueTitle;
                            this.userKey = json.UserKey;
                            this.queueState = json.QueueState;
                            if (this.userPulledNumber > 0 && this.userPulledNumber <= this.queueState.LastCalled)
                            {
                                this.showOwnNumberWasCalled = true
                            }
                        } else {
                            this.showTechnicalError = true;
                            setTimeout(() => this.showTechnicalError=false, 5000);
                        }
                    });
            },
            pullNext: function() {
                fetch(process.env.VUE_APP_URL + '/user/' + this.userKey + "/pull")
                    .then(response => {
                        if (response.status === 200) {
                            return response.json();
                        } else {
                            console.warn('response failed');
                            return {};
                        }
                    })
                    .catch(() => {
                        console.warn("Network connection error");
                        return {};
                    })
                    .then(json => {
                        if (json.QueueTitle) {
                            this.queueTitle = json.QueueTitle;
                            this.queueState = json.QueueState;
                            this.userPulledNumber = json.QueueState.LastPulled;
                            this.userCanPull = false;
                            // save as cookie:
                            this.$cookies.set(this.userKey, this.userPulledNumber);
                        } else {
                            this.showTechnicalError = true;
                            setTimeout(() => this.showTechnicalError=false, 5000);
                        }
                    });
            },
            createUserUrl: function() {
                return window.location.toString().replace(/#.*/, "#/user/" + this.userKey);
            },
            createBaseUrl: function() {
                return window.location.toString().replace(/#.*/, "");
            }
        },
        beforeDestroy () {
            clearInterval(this.timer)
        }
    }
</script>

<style scoped>

    .queueTitle
    {
        font-size: 200%;
    }

    .userPulledNumber
    {
        font-size: 400%;
    }

    .showOwnNumberWasCalled
    {
        font-size: 400%;
    }

    .lastCalled
    {
        font-size: 400%;
    }

    .userIsCalled {
        text-align: center;
    }
    
    .userKey
    {
        font-family: monospace;
        color: blue;
    }

    .baseUrl
    {
        font-family: monospace;
        color: blue;
    }

</style>