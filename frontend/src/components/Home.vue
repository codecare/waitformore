<template>
    <v-container>

        <v-row class="queueUpRow">
            <v-col class="shrink">
                <img src="../assets/user.png" :alt="$t('altUserImage')" width="150" height="150"/>
            </v-col>
            <v-col class="grow">

                <div v-if="$i18n.locale === 'de'" >
                    <h3>Ich will mich per Handy anstellen</h3>
                    <p>und habe den Code einer Warteschlange erhalten.</p>
                    <p>Bitte geben Sie den 4-Buchstaben Code der Warteschlange ein.</p>
                </div>

                <div v-if="$i18n.locale === 'en'" >
                    <h3>I would like to join the queue</h3>
                    <p>and I have my 4-letter code at the ready.</p>
                    <p>Please enter the 4-letter of the queue.</p>
                </div>

                <v-form ref="form" v-model="userKeyValid" @submit="queueUpValid">
                    <v-text-field
                            :label="$t('queueUpLabel')"
                            v-model="userKey"
                            :rules="[() => !!userKey || $t('pleaseFillIn'), v => (v && v.length === 4) || $t('codeHas4Letters')]"
                            required
                            counter=4
                    ></v-text-field>
                    <v-btn style="margin:5px;" @click="queueUp" :disabled="!userKeyValid">{{ $t( 'queueUpButton' ) }}</v-btn>
                    <v-btn style="margin:5px;" href="#usage">{{ $t( 'navToUsage' )}}</v-btn>
                </v-form>

                <v-spacer>&nbsp;</v-spacer>
                <v-alert type="warning"
                         dismissible
                         transition="scale-transition"
                         :value="showQueueNotFound">
                    {{ $t('message.queueNotFound', { msg: userKey })}}
                </v-alert>
                <v-alert type="warning"
                         dismissible
                         transition="scale-transition"
                         :value="showTechnicalError">
                    {{ $t('technicalProblem') }}
                </v-alert>
            </v-col>
        </v-row>

        <v-row>
            <v-col>
                <v-divider></v-divider>
            </v-col>
        </v-row>

        <v-row class="provideQueueRow">
            <v-col class="shrink">
                <img src="../assets/queueAdmin.png" :alt="$t('altProvideQueue')" width="150" height="150"/>
            </v-col>
            <v-col class="grow">
                <div v-if="$i18n.locale === 'de'">
                    <h3>Ich will eine Warteschlange bereitstellen</h3>
                    <p>Erstellen Sie hier eine neue Warteschlange für Ihre Kunden.</p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <h3>I would like to create a queue for others to join</h3>
                    <p>Create a new queue for customers here.</p>
                </div>
                <v-btn to="createQueue">{{ $t('provideQueueButton')}}</v-btn>
            </v-col>
        </v-row>


        <v-row>
            <v-col>
                <v-divider></v-divider>
            </v-col>
        </v-row>

        <v-row>
            <v-col>
                <h3 id="usage">{{ $t('usage') }}</h3>
                <br>
                <div v-if="$i18n.locale === 'de'">
                    <h4>Zum Anstellen:</h4>
                    <br>
                    <p>Sie erhalten den Zugriff zur Warteschlange von dem Verwalter,
                        entweder als QR-Code zum Scannen oder als vierstelligen Buchstabencode zum Eingeben.
                        <br>Dadurch kommen Sie zur Website Ihrer gewünschten Warteschlange.</p>
                    <p>Dann einfach "Wartenummer ziehen" drücken und warten, bis Sie von der App und vom Personal aufgerufen werden.</p>
                    <p>Sie können live verfolgen, wie viele Wartenummern vor Ihnen in der Schlange sind.</p>
                    <p>Bei Problemen bitte dem Warteschlangenverwalter Bescheid geben - er kann für Sie eine Nummer ziehen. </p>
                    <br>

                    <h4>Zum Warteschlangen Bereitstellen und Verwalten</h4>
                    <br>
                    <p>Benutzen Sie den Button oben, um eine neue Warteschlange zu erstellen.
                        Diese erhält einen Link, mit dem Sie und andere Mitarbeiter Kunden aufrufen können.</p>
                    <p>Geben Sie den QR-Code und die Website mit dem vierstelligen Buchstabencode an Kunden weiter,
                        damit sie sich anstellen können.
                        <br> Wir haben einen PDF-Ausdruck mit dieser Information für Sie bereitgestellt - bitte finden Sie ihn unter "PDF generieren" in der Warteschlangenverwaltung.
                        <br>Kunden ohne Zugriff auf die App können von Ihnen in die Warteschlange eingefügt werden - bitte teilen Sie ihnen ihre Nummer mit.</p>
                    <p>Rufen Sie ganz einfach die nächsten Kunden auf.
                        Die Kunden mit der App können neben ihrer eigenen Nummer sehen, welche Nummer zuletzt gerufen wurde.</p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <h4>To join the queue:</h4>
                    <br>
                    <p>You will receive access to the queue from the queue administrator,
                        either as QR-Code or as a 4-letter code.
                        <br>You can use either to find the website of your queue.</p>
                    <p>Then, simply press "Take a number" and wait for your turn.</p>
                    <p>You can follow your position in the queue in real time.</p>
                    <p>If you are having trouble please contact the queue administrator - they can take a number for you.</p>
                    <br>

                    <h4>To create a queue and manage it:</h4>
                    <br>
                    <p>Please use the button above to create a new queue.
                        <br>You will be directed to the queue administration, where you can call customers when its their turn.
                        Your queue is assigned an administration link, which you can distribute amongst fellow queue administrators, 
                        so multiple people can call customers from the same queue.</p>
                    <p>Please distribute the QR-Code and the website with the 4-letter code to customers wanting to join your queue.
                        <br> We have prepared a PDF-page for printing with everything customers need to know - you can find it under "generate PDF" on the queue administration page.
                        <br>Customers without access to the app or a phone can be added to the queue by the administrators - please tell them their number in the queue.</p>
                    <p>Simply call the next customer when you are ready.
                        Customers with the app can follow their position in the queue, and will be notified when it is their turn.</p>
                </div>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <div v-if="$i18n.locale === 'de'">
                    <h3>Weiterentwicklung</h3>
                    <br>
                    <p>Wir sind aktiv dabei, diese Anwendung weiterzuentwickeln.</p>
                    <p>Aktuell arbeiten wir an:</p>
                    <ul>
                        <li>Übersetzung in Englisch, Spanisch und Türkisch</li>
                        <li>Einem Info-Screen mit dem Fortschritt der Schlange</li>
                    </ul>
                    <br>
                    <p>Fragen gerne an info@wartewarte.de</p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <h3>Website development</h3>
                    <br>
                    <p>We are making this app better every day.</p>
                    <p>Our next goals are:</p>
                    <ul>
                        <li>Language support in English, Spanish and Turkish</li>
                        <li>A link to an information screen which can be displayed to customers</li>
                    </ul>
                    <br>
                    <p>We would love to hear your feedback and ideas: info@wartewarte.de</p>
                </div>
            </v-col>
        </v-row>
        <v-row>
            <v-col>

                <div class="version">{{ $t('versionInfo' )}} {{ $t('version') }}</div>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>


    export default {
        data: () => ({
            queueName: '',
            userKeyValid: true,
            formHasErrors: false,
            userKey: null,
            showQueueNotFound: false,
            showTechnicalError: false
        }),
        created() {
            console.info("backend location set to " + process.env.VUE_APP_URL  );
        },

        methods: {
            queueUpValid: function() {
                if (this.userKeyValid) {
                    this.queueUp();
                }
            },
            queueUp: function() {

                fetch(process.env.VUE_APP_URL + '/user/' + this.userKey)
                    .then(response => {
                        if (response.status === 200) {
                            return 200;
                        } else if (response.status === 404) {
                            return 404;
                        } else {
                            console.warn('response failed ' + response.status);
                            return 500;
                        }
                    })
                    .catch(() => {
                        (console.warn("Network connection error"));
                        return 500;
                    })
                    .then(status => {
                        if (status === 200) {
                            this.$router.push({ path: `/user/${this.userKey}` });
                        } else if (status === 404) {
                            this.showQueueNotFound = true;
                            console.warn("queue not found by userKey: " + this.userKey)
                            setTimeout(() => this.showQueueNotFound=false, 5000);
                        } else if (status === 500) {
                            this.showTechnicalError = true;
                            setTimeout(() => this.showTechnicalError=false, 5000);
                        }
                    });
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    .queueUpRow {
        border-radius: 4px;
        border-color: #4CAF50;
        border-style: solid;
    }

    .provideQueueRow {
        border-radius: 4px;
        border-color: #2196F3;
        border-style: solid;
    }
    a {
        color: #42b983;
    }
</style>
