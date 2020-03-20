import Vue from 'vue';
import VueI18n from 'vue-i18n'

import { version } from "../../package.json";

Vue.use(VueI18n);

const messages = {
    de: {
        appUrl: "wartewarte.de",

        versionNumber: version,
        versionLabel: 'Frontend Versionsnummer: ',
        headerText: 'Wartenummern auf dem Handy',
        altUserImage: 'an Schlange anstellen',
        queueUpLabel: '4-Buchstaben-Code. Klein-/Großschreibung beachten',
        queueUpButton: 'Anstellen',
        pleaseFillIn: 'Bitte ausfüllen',
        codeHas4Letters: 'Der Code hat exakt 4 Zeichen',
        navToUsage: 'Zur Anleitung',

        technicalProblem: 'Entschuldigung - technisches Problem!',
        altProvideQueue: 'Warteschlange bereitstellen',
        provideQueueButton: 'Bereitstellen',
        usage: 'Anleitung',

        queueName: 'Name der Warteschlange',
        pleaseFillInQueueName: 'Bitte Warteschlangen Namen eingeben',
        nameHasMax30Letters: 'Der Name hat maximal 30 Zeichen',
        createQueueButton: 'Neue Warteschlange erstellen',

        privacyPolicy: 'Datenschutzerklärung',
        imprintTitle: 'Lizenshinweis und Impressum',
        cookiePolicy: 'Cookie Policy',

        queueAdmin: 'Verwaltung der Warteschlange: ',
        altLastCalled: 'zuletzt aufgerufen',
        lastCalled: 'zuletzt aufgerufen',
        buttonCallNext: 'Nächsten aufrufen',
        nobodyWaitingOnCall: 'Niemand wartet gerade!',
        altLastPulled: 'zuletzt gezogen',
        lastPulled: 'zuletzt gezogen',
        pullForCustomer: 'Für Kunden ziehen',

        btnRefresh: 'Seite neu laden',
        autoRefresh: 'Die Seite aktualisiert sich automatisch alle 10 Sekunden',

        userQueueName: 'Warteschlange: ',
        userLastCalled: 'zuletzt aufgerufen',
        altQueueUp: 'an Schlange anstellen',

        userIsCalled: 'Sie wurden aufgerufen!',
        userWaitingNumber: 'Meine Wartenummer',
        pullWaitingNumber: 'Wartenummer ziehen',

        hasPhoneNoQr: "Sie haben ein Handy ohne QR Scanner?",
        hasQr: "Sie haben einen QR Scanner?",
        hasNoPhone: "Sie haben kein Handy?",

        queueInstructionsHeader: "Bitte beachten Sie folgendes, um sich anzustellen:",
        navigateToPage: "Bitte navigieren Sie im Browser zu",
        enterThisCode: "und geben Sie diesen Code ein:",
        isCaseSensitive: "Auf Klein-/Großschreibung achten!",
        clickTakeNumber: "Drücken Sie dann \"Wartenummer Ziehen\", um sich anzustellen.",
        pleaseScanQr: "Bitte scannen Sie diesen QR-Code.",
        talkToAdmin: "Bitte geben Sie dem Personal Bescheid.",
        numberWillBeAssigned: "Es kann eine Nummer für Sie gezogen, und Ihnen mitgeteilt werden.",
        pdfSaveName: "anstellen_anleitung.pdf",

        message: {
            queueNotFound: 'Warteschlange {msg} nicht gefunden!',
            lastPulledForCustomer: 'Zuletzt für einen Kunden gezogen {msg}',
        }

    },
    en: {
        appUrl: "wartewarte.de",

        versionNumber: version,
        versionLabel: 'Frontend version: ',
        headerText: 'Smart queueing',
        altUserImage: 'join queue',
        queueUpLabel: '4 letter code (case-sensitive)',
        queueUpButton: 'Join queue',
        pleaseFillIn: 'Please ener your code to proceed',
        codeHas4Letters: 'The code has 4 letters exactly',
        navToUsage: 'How does this work?',
        technicalProblem: 'We are sorry - there seems to be a technical error',
        altProvideQueue: 'Create a new queue',
        provideQueueButton: 'Create queue',
        usage: 'How to use this app',

        queueName: 'Name of the queue',
        pleaseFillInQueueName: 'Please enter a name for your queue',
        nameHasMax30Letters: 'Please use 30 letters or less',
        createQueueButton: 'Create new queue',

        privacyPolicy: 'Data and Privacy Policy',
        imprintTitle: 'Legal notice',
        cookiePolicy: 'Cookie Policy',

        queueAdmin: 'Queue administration: ',
        altLastCalled: 'last called',
        lastCalled: 'last called',
        buttonCallNext: 'Call next customer',
        nobodyWaitingOnCall: 'Nobody is waiting!',
        altLastPulled: 'last number taken',
        lastPulled: 'last number taken',
        pullForCustomer: 'take number for customer',

        btnRefresh: 'Refresh page',
        autoRefresh: 'This page refreshes every 10 seconds',

        userQueueName: 'Queue: ',
        userLastCalled: 'last called',
        altQueueUp: 'join this queue',

        userIsCalled: 'It\'s your turn',
        userWaitingNumber: 'my number in the queue',
        pullWaitingNumber: 'Take a number',
        
        hasPhoneNoQr: "If you have a smartphone without QR scanner",
        hasQr: "If you have a QR scanner",
        hasNoPhone: "If you don't have a smartphone",

        queueInstructionsHeader: "Please take note of the following to join the queue:",
        navigateToPage: "In your browser, please navigate to",
        enterThisCode: "and enter the following code:",
        isCaseSensitive: "The code is case-sensitive!",
        clickTakeNumber: "Then press the button \"Take a number\" to join the queue.",
        pleaseScanQr: "Please scan the QR code below",
        talkToAdmin: "Please notify a queue administrator.",
        numberWillBeAssigned: "They can make sure you are queued and can inform you of your number.",
        pdfSaveName: "queue_instructions.pdf",

        message: {
            queueNotFound: 'The queue {msg} could not be found!',
            lastPulledForCustomer: 'Last number taken for a customer: {msg}',
        }
    }
};

// Create VueI18n instance with options
const i18n = new VueI18n({
    locale: 'de', // set locale
    fallbackLocale: 'de',
    messages, // set locale messages
});

export default i18n;