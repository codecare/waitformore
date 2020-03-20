<template>
    <v-container>

        <v-row>
            <v-col>
                <h3>{{ $t('queueAdmin') }} <br>
                    <span class="queueTitle"> {{ queueTitle }}</span></h3>
                <div v-if="$i18n.locale === 'de'">
                    <p>Jeder der möchte, kann diesen Service kostenlos verwenden.<br>
                        Es ist jedoch eine Beta-Version - wir übernehmen keine Haftung oder Verantwortung bei Fehlern oder Ausfällen!</p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <p>This service is freely available to everyone.<br>
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
                    {{ $t('technicalProblem' )}}
                </v-alert>

                <v-alert color="blue" dark>
                    <v-row align="center">
                        <v-col class="shrink">
                            <img src="../assets/queueAdmin-white.png" :alt="$t('altLastCalled')" width="150" height="150"/>
                        </v-col>
                        <v-col class="grow"><span class="lastCalled">{{ queueState.LastCalled }}</span> {{ $t('lastCalled')}}</v-col>
                        <v-col class="shrink">
                            <v-btn @click="callNext">{{ $t('buttonCallNext' )}}</v-btn>
                        </v-col>
                    </v-row>
                </v-alert>

                <v-alert type="warning"
                         dismissible
                         transition="scale-transition"
                         :value="showNobodyWaiting">
                    {{ $t('nobodyWaitingOnCall') }}
                </v-alert>

            </v-col>
        </v-row>

        <v-row>
            <v-col>

                <v-alert color="green" dark>
                    <v-row align="center">
                        <v-col class="shrink">
                            <img src="../assets/user-white.png" :alt="$t('lastPulled')" width="150" height="150"/>
                        </v-col>
                        <v-col class="grow"><span class="lastPulled">{{ queueState.LastPulled }}</span> {{ $t('lastPulled') }}</v-col>
                        <v-col class="shrink">
                            <v-btn @click="pullNext">{{ $t('pullForCustomer') }}</v-btn>
                        </v-col>
                    </v-row>
                </v-alert>

                <v-alert type="warning"
                         dismissible
                         transition="scale-transition"
                         :value="showPulledForCustomer">
                    {{ $t('message.lastPulledForCustomer', { msg: LastPulledForCustomer })}}
                </v-alert>

            </v-col>
        </v-row>



        <v-row>
            <v-col>
                <div v-if="$i18n.locale === 'de'">
                    <h4>Name Ihrer Warteschlange: {{ queueTitle }}</h4>
                    <h3 ><br>Hier erhalten Sie eine ausdruckbare Anleitung für Nutzer Ihrer Warteschlange</h3>
                    <v-btn href="#" @click="generatePDF" style="margin:10px;">PDF generieren </v-btn>
                    <p><br>Kunden-Link:
                        <br><a :href="userUrl" class="baseUrl">{{ userUrl }}</a></p>
                <p>Alternativ kann der Code <span class="userKey">{{ userKey }}</span> auf der Homepage <a :href="baseUrl" class="baseUrl">{{ baseUrl }}</a> eingegeben werden.</p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <h3>Link to queue: {{ queueTitle }}</h3>
                    <h3 ><br>Press the button below to save and print instructions for your customers</h3>
                    <v-btn href="#" @click="generatePDF" style="margin:10px;">Generate PDF</v-btn>
                    <p><br>Link for customers
                        <br><a :href="userUrl" class="baseUrl">{{ userUrl }}</a></p>
                    <p>You may also tell them to enter the code <span class="userKey">{{ userKey }}</span> on the home page <a :href="baseUrl" class="baseUrl">{{ baseUrl }}</a></p>
                </div>
            </v-col>
            <v-col>
                <a :href="userUrl" :target="userKey" class="baseUrl">
                    <qrcode-vue :value="userUrl" size="300" level="H" id="userQrCode"></qrcode-vue>
                    {{ userUrl }}
                </a>
            </v-col>
        </v-row>

        <v-row>
            <v-col>
                <div v-if="$i18n.locale === 'de'">
                    <h3>Url zur Verwaltung der Warteschlange: {{ queueTitle }}</h3>
                    <p>Nur an Verwalter weitergeben.</p>
                    <p><a  class="userKey" :href="adminUrl">{{ adminUrl }}</a></p>
                </div>
                <div v-if="$i18n.locale === 'en'">
                    <h3>Link for queue administrators: {{ queueTitle }}</h3>
                    <p>Please only distribute to people who need to call customers from this queue.</p>
                    <p><a class="userKey" :href="adminUrl">{{ adminUrl }}</a></p>
                </div>
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

    import QrcodeVue from 'qrcode.vue'
    import jsPDF from 'jspdf'

    export default {
        name: "QueueAdmin",
        data: () => ({
            queueTitle: '',
            adminKey: '',
            userKey: '',
            queueState: {},
            baseUrl: '',
            userUrl: '',
            adminUrl: '',
            showNobodyWaiting: false,
            timer: '',
            showPulledForCustomer: false,
            LastPulledForCustomer: '',
            showTechnicalError: false
        }),
        components: {
            QrcodeVue
        },
        created() {
            this.adminKey = this.$route.params.adminKey;
            this.baseUrl = this.createBaseUrl();
            this.adminUrl = this.createAdminUrl();
            this.refresh();
            this.timer = setInterval(this.refresh, 10000);
        },
        methods: {
            refresh: function() {
                fetch(process.env.VUE_APP_URL + '/queue/' + this.adminKey)
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
                        if (json.UserKey) {
                            this.queueTitle = json.QueueTitle;
                            this.userKey = json.UserKey;
                            this.queueState = json.QueueState;
                            this.userUrl = this.createUserUrl();
                        } else {
                            this.showTechnicalError = true;
                            setTimeout(() => this.showTechnicalError=false, 5000);
                        }
                    });
            },
            callNext: function() {
                fetch(process.env.VUE_APP_URL + '/queue/' + this.adminKey + "/call")
                    .then(response => {
                        if (response.status === 200) {
                            response.json().then(json => {
                                    this.queueState = json.QueueState;
                                    this.queueTitle = json.QueueTitle;
                                }
                            );
                        } else if (response.status === 406) {
                            this.showNobodyWaiting = true;
                            setTimeout(() => this.showNobodyWaiting=false, 2000);
                        } else {
                            this.showTechnicalError = true;
                            setTimeout(() => this.showTechnicalError=false, 5000);
                        }
                    })
                    .catch(() => {
                        this.showTechnicalError = true;
                        setTimeout(() => this.showTechnicalError=false, 5000);
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
                        (console.warn("Network connection error"));
                        return {};
                    })
                    .then(json => {
                        if (json.UserKey) {
                            this.queueTitle = json.QueueTitle;
                            this.queueState = json.QueueState;
                            this.LastPulledForCustomer = json.QueueState.LastPulled;
                            this.showPulledForCustomer = true;
                        } else {
                            this.showTechnicalError = true;
                            setTimeout(() => this.showTechnicalError=false, 5000);
                        }
                    });
            },
            createUserUrl: function() {
                return window.location.toString().replace(/#.*/, "#/user/" + this.userKey);
            },
            createAdminUrl: function() {
                return window.location.toString().replace(/#.*/, "#/queueAdmin/" + this.adminKey);
            },
            createBaseUrl: function() {
                return window.location.toString().replace(/#.*/, "");
            },
            generatePdfHeader: function(doc, pageWidth, pageHozMiddle) {
                doc.setTextColor(0, 0, 0);
                doc.setFontSize("24");
                doc.setFontStyle("bold");
                doc.text(this.queueTitle, pageHozMiddle, 30, {"align":"center", "maxWidth": pageWidth-20});
                doc.setFontSize("35");
                doc.setFontStyle("bold");
                doc.text(this.$t('queueInstructionsHeader'), pageHozMiddle, 55, {"align":"center", "maxWidth": pageWidth-20});
            },
            generatePDF: function() {
                let doc = new jsPDF();
                const pageHeight = doc.internal.pageSize.height || doc.internal.pageSize.getHeight();
                const pageWidth = doc.internal.pageSize.width || doc.internal.pageSize.getWidth();
                const pageHozMiddle = pageWidth / 2;
                const pageVertMiddle = pageHeight / 2;            
                const defaultFont = "helvetica";
                const usrKeyFont = "courier"

                doc.setFont(defaultFont);
                
                /******************************/
                /* First Page (No QR-Scanner) */
                /******************************/
                this.generatePdfHeader(doc, pageWidth, pageHozMiddle);

                // Orange Background
                doc.setDrawColor(251, 140, 0);
                doc.setFillColor(251, 140, 0);
                doc.roundedRect(pageWidth * 0.05, pageHeight * 0.265, pageWidth * 0.9, pageHeight * 0.125, 5, 5, 'F');
                // Title
                doc.setDrawColor(0, 0, 0)
                doc.setFontStyle("normal");
                doc.setFontSize("35");
                doc.text(this.$t('hasPhoneNoQr'), pageHozMiddle, 95, {"align":"center", "maxWidth": pageWidth-30}); //margin larger for text in boxes
                // Text
                doc.setFontSize("24");
                doc.text(this.$t('navigateToPage'), pageHozMiddle, 135, {"align":"center", "maxWidth": pageWidth-20});
                doc.setFontSize("55");
                doc.text(this.$t('appUrl'), pageHozMiddle, 155, {"align":"center", "maxWidth": pageWidth-20});
                doc.setFontSize("24");
                doc.text(this.$t('enterThisCode'), pageHozMiddle, 170, {"align":"center", "maxWidth": pageWidth-20});
                // Code
                doc.setFontSize("120");
                doc.setFont(usrKeyFont);
                doc.text(this.userKey, pageHozMiddle, 215, {"align":"center", "maxWidth": pageWidth-20});
                doc.setFont(defaultFont);
                doc.setFontSize("24");
                doc.setTextColor(255, 0, 0);
                doc.text(this.$t('isCaseSensitive'), pageHozMiddle, 235, {"align":"center", "maxWidth": pageWidth-20});
                // Green Background
                doc.setDrawColor(76, 175, 80);
                doc.setFillColor(76, 175, 80);
                doc.roundedRect(pageWidth * 0.05, pageHeight * 0.82, pageWidth * 0.9, pageHeight * 0.11, 5, 5, 'F');
                // Footer
                doc.setTextColor(255, 255, 255);
                doc.setFontSize("24");
                doc.text(this.$t('clickTakeNumber'), pageHozMiddle, pageHeight * 0.86, {"align":"center", "maxWidth": pageWidth-30});

                
                /****************************/
                /* Second Page (QR-Scanner) */
                /****************************/
                doc.addPage();
                this.generatePdfHeader(doc, pageWidth, pageHozMiddle);

                // Orange Background
                doc.setDrawColor(251, 140, 0);
                doc.setFillColor(251, 140, 0);
                doc.roundedRect(pageWidth * 0.05, pageHeight * 0.265, pageWidth * 0.9, pageHeight * 0.125, 5, 5, 'F');
                // Title
                doc.setDrawColor(0, 0, 0)
                doc.setFontStyle("normal");
                doc.setFontSize("35");
                doc.text(this.$t('hasQr'), pageHozMiddle, 100, {"align":"center", "maxWidth": pageWidth-30});
                // Text
                doc.setFontSize("28");
                doc.text(this.$t('pleaseScanQr'), pageHozMiddle, 135, {"align":"center", "maxWidth": pageWidth-20});
                // QR Code
                let qrCodeCanvas = document.getElementById("userQrCode").querySelector("canvas");
                doc.addImage(qrCodeCanvas.toDataURL(), pageHozMiddle-40, pageVertMiddle, 80, 80);
                doc.setFontSize("16");
                doc.setFont(usrKeyFont);
                doc.text(this.userUrl, pageHozMiddle, pageVertMiddle + 90, {"align":"center", "maxWidth": pageWidth-20})
                doc.setFont(defaultFont);
                // Green Background
                doc.setDrawColor(76, 175, 80);
                doc.setFillColor(76, 175, 80);
                doc.roundedRect(pageWidth * 0.05, pageHeight * 0.82, pageWidth * 0.9, pageHeight * 0.11, 5, 5, 'F');
                // Footer
                doc.setTextColor(255, 255, 255);
                doc.setFontSize("24");
                doc.text(this.$t('clickTakeNumber'), pageHozMiddle, pageHeight * 0.86, {"align":"center", "maxWidth": pageWidth-30});
                
                /******************************/
                /* Third Page (No Smartphone) */
                /******************************/
                //No smartphone
                doc.addPage();  
                this.generatePdfHeader(doc, pageWidth, pageHozMiddle);

                // Orange Background
                doc.setDrawColor(251, 140, 0);
                doc.setFillColor(251, 140, 0);
                doc.roundedRect(pageWidth * 0.05, pageHeight * 0.265, pageWidth * 0.9, pageHeight * 0.125, 5, 5, 'F');
                // Title
                doc.setDrawColor(0, 0, 0)
                doc.setFontStyle("normal");
                doc.setFontSize("35");
                doc.text(this.$t('hasNoPhone'), pageHozMiddle, 100, {"align":"center", "maxWidth": pageWidth-30});
                // Text
                doc.setFontSize("28");
                doc.text(this.$t('talkToAdmin'), pageHozMiddle, 140, {"align":"center", "maxWidth": pageWidth-20});
                doc.text(this.$t('numberWillBeAssigned'), pageHozMiddle, 160, {"align":"center", "maxWidth": pageWidth-20});



                //doc.save('queue.pdf');
                doc.save(this.$t('pdfSaveName'));

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

    .lastCalled
    {
        font-size: 400%;
    }

    .lastPulled
    {
        font-size: 400%;
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