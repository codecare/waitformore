<template>
    <v-footer>
        <v-row>
            <v-col>
                <v-expansion-panels>
                    <v-expansion-panel>
                        <v-expansion-panel-header>{{ $t('privacyPolicy')}}</v-expansion-panel-header>
                        <v-expansion-panel-content>
                            <div v-if="$i18n.locale === 'de'">
                                <h3>Rechtsgrundlagen der Datenverarbeitung</h3>
                                <p>Durch die Nutzung dieser Website erklären Sie sich mit der Erhebung, Nutzung und Löschung von Daten
                                    gemäß der folgenden Beschreibung einverstanden.
                                    Diese Website kann grundsätzlich ohne Registrierung besucht werden. Dabei werden Daten wie
                                    IP-Adresse des Aufrufers und aufgerufene Seiten protokolliert.
                                    Dies geschieht zu statistischen Zwecken und zur Ermöglichung der Analyse technischer Fehler seitens
                                    des Betreibers.</p>

                                <p>Namen der abgerufenen Datei, Datum und Uhrzeit werden zu statistischen Zwecken auf dem Server
                                    gespeichert, ohne dass diese Daten unmittelbar auf Ihre Person bezogen werden.
                                    Diese Daten werden spätestens nach zwei Wochen automatisch gelöscht.</p>

                                <p>Soweit auf unseren Seiten personenbezogene Daten (beispielsweise Name, Anschrift oder
                                    eMail-Adressen) erhoben werden, erfolgt dies, soweit möglich, stets auf freiwilliger Basis.
                                    Ohne Ihre Einwilligung erfolgt keine Weitergabe der Daten an Dritte.
                                    Die Nutzung unserer Webseite ist ohne Angabe personenbezogener Daten möglich. </p>

                                <h3>Nutzung der Impressumsdaten durch Dritte</h3>
                                <p>Der Nutzung von im Rahmen der Impressumspflicht veröffentlichten Kontaktdaten
                                    durch Dritte zur Übersendung von nicht ausdrücklich angeforderter
                                    Werbung und Informationsmaterialien wird hiermit ausdrücklich widersprochen.
                                    Die Betreiber der Seiten behalten sich ausdrücklich rechtliche Schritte
                                    im Falle der nicht verlangten Zusendung von Werbeinformationen, etwa durch Spam-Mails, vor.</p>
                            </div>
                            <div v-if="$i18n.locale === 'en'">
                                <h3>Data Policy</h3>
                                <p>By using this website you consent to the collection, use and deletion of data subject to the following conditions.</p>
                                <p>This website can be used without registering. We do collect data such as
                                    IP-addresses of the visitor and which websites have been visited.
                                    We use this data for statistic analysis and for correction of technical faults.</p>

                                <p>The names of visited websites, time and date are collected by the server for statistical analysis purposes, 
                                    without any direct association to personal data.
                                    Such data is deleted automatically after two weeks.</p>

                                <p>This website is fully functional without the collection of personal data of its users. 
                                    Any collection of personal data (such as name, address or e-mail addresses)
                                    is optional.
                                    Without your explicit consent we will not forward any of your data to third parties.
                                    </p>

                                <h3>Use of the legal notice data provided by us</h3>
                                <p>We strongly oppose the misuse of the data we provide in our legal notice for not explicitly requested information and advertising. 
                                    We reserve the right to take legal action against unwarranted messages, such as advertising and spam-emails.</p>
                            </div>
                        </v-expansion-panel-content>
                    </v-expansion-panel>
                </v-expansion-panels>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-expansion-panels>

                    <v-expansion-panel>
                        <v-expansion-panel-header>{{ $t('imprintTitle') }}</v-expansion-panel-header>
                        <v-expansion-panel-content>
                            <div v-if="$i18n.locale === 'de'">
                                <h3>Hinweis / Lizenz</h3>
                                <p>Jeder der möchte, kann diesen Service kostenlos verwenden.
                                    <br>Es ist jedoch eine Beta-Version - wir übernehmen keine Haftung oder Verantwortung bei Fehlern oder Ausfällen!</p>
                                <h3>Betreiber der Seite</h3>
                                <p>codecare software GmbH<br>
                                    Moosbergstraße 42<br>
                                    64285 Darmstadt</p>

                                <h3>Datenschutzbeauftragter</h3>
                                <p>Thorsten Rottschäfer<br>
                                    E-Mail: tr@codecare.de<br>
                                    Phone: +49 176 84 26 53 50</p>

                                <h3>Auskunftsrecht</h3>
                                <p>Wenn Sie Fragen hinsichtlich der Verarbeitung Ihrer persönlichen Daten haben, können Sie sich direkt
                                    an unseren Beauftragten für den Datenschutz wenden.
                                    <br>
                                    Er steht Ihnen auch im Falle von Auskunftsersuchen,
                                    Anträgen oder Beschwerden zur Verfügung.</p>
                            </div>
                            <div v-if="$i18n.locale === 'en'">
                                <h3>Legal notice</h3>
                                <p>This service is freely available to everyone.
                                    <br>This is a Beta-version - we are very invested in making sure it all works, however we cannot be held responsible for failures.</p>
                                <h3>This website is maintained by</h3>
                                <p>codecare software GmbH<br>
                                    Moosbergstraße 42<br>
                                    64285 Darmstadt</p>

                                <h3>Data protection officer</h3>
                                <p>Thorsten Rottschäfer<br>
                                    E-Mail: tr@codecare.de<br>
                                    Phone: +49 176 84 26 53 50</p>

                                <h3>Right to information</h3>
                                <p>Should you have questions regarding your personal data 
                                    you can direct these to our data protection officer, who will receive your requests for information and any complaints.</p>
                            </div>
                        </v-expansion-panel-content>
                    </v-expansion-panel>
                </v-expansion-panels>
            </v-col>
        </v-row>
        <v-row>
            <v-col>

                <v-expansion-panels>
                    <v-expansion-panel>
                        <v-expansion-panel-header>Cookie Policy</v-expansion-panel-header>
                        <v-expansion-panel-content>
                            <div v-if="$i18n.locale === 'de'">
                                <h3>Cookie(s)</h3>
                                <p> Diese Seite verwendet genau ein technisch notwendiges Cookie pro Benutzer.
                                    <br>Wir benutzen dieses, damit wir Ihre Warteschlangenposition wieder aufrufen können.</p>
                            </div>
                            <div v-if="$i18n.locale === 'en'">
                                <h3>Cookie(s)</h3>
                                <p> This web page requires precisely one cookie for users.
                                    <br>We use this cookie so customers are able to keep their position in the queue.</p>
                            </div>
                        </v-expansion-panel-content>
                    </v-expansion-panel>
                </v-expansion-panels>
            </v-col>
        </v-row>
    </v-footer>
</template>

<script>
</script>

<style scoped>
</style>