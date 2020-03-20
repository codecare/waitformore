<template>
    <v-container>

        <v-row>
            <v-col>
                <div v-if="$i18n.locale === 'de'">
                    <h3>Neue Warteschlange</h3>
                    <p>Jeder der möchte, kann diesen Service kostenlos verwenden.
                        <br>Es ist jedoch eine Beta-Version - wir übernehmen keine Haftung oder Verantwortung bei Fehlern oder Ausfällen!</p>
                </div>
                <div v-if="$i18n.locale === 'end'">
                    <h3>Create a new queue</h3>
                    <p>This service is freely available to everyone.
                        <br>This is a Beta-version - we are very invested in making sure it all works, however we cannot be held responsible for failures.</p>
                </div>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-form ref="form" v-model="valid" @submit="createQueueValid">

                    <v-text-field
                            :label="$t('queueName')"
                            ref="queueTitle"
                            v-model="queueTitle"
                            :rules="[() => !!queueTitle || $t('pleaseFillInQueueName'), v => (v && v.length <= 30) || $t('nameHasMax30Letters')]"
                            required
                            counter=30></v-text-field>
                    <v-btn @click="createQueue" :disabled="!valid">{{ $t('createQueueButton')}}</v-btn>
                </v-form>
                <v-spacer>&nbsp;</v-spacer>
                <v-alert type="warning"
                         dismissible
                         transition="scale-transition"
                         :value="showTechnicalError">
                    {{ $t('technicalProblem') }}
                </v-alert>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>

    export default {
        name: "CreateQueue",
        data: () => ({
            queueTitle: '',
            valid: true,
            showTechnicalError: false
        }),
        methods: {
            createQueueValid: function() {
                if (this.valid) {
                    this.createQueue();
                }
            },
            createQueue: function() {
                fetch(process.env.VUE_APP_URL + '/queue', {
                    method: 'post',
                    body: JSON.stringify({queueTitle: this.queueTitle})
                })
                    .then(response => {

                        if (response.status === 200) {
                            response.json().then(json => {
                                this.adminKey = json.AdminKey;
                                this.$router.push({path: `/queueAdmin/${this.adminKey}`});
                            });
                        } else {
                            console.warn('response failed');
                            this.showTechnicalError = true;
                            setTimeout(() => this.showTechnicalError=false, 5000);
                        }
                    })
                    .catch(() => {
                        console.warn("Network connection error");
                        this.showTechnicalError = true;
                        setTimeout(() => this.showTechnicalError=false, 5000);
                        return 400;
                    });
            }
        }
    }
</script>

<style scoped>

</style>