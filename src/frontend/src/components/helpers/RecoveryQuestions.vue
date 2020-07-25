<template>
  <span>
    <span v-for="(v,i) in sets" :key="i">
      <v-row>
        <v-select
          v-model="sets[i].id"
          :items="questions"
          :label="'Question #'+(i+1)"
          :rules="[$validator.required,uniqueQuestion(i, sets)]"
          :disabled="predefined"
        />
      </v-row>
      <v-row>
        <v-text-field
          v-model="sets[i].answer"
          label="Answer"
          type="password"
          class="purple-input"
          :rules="[$validator.required]"
        />
      </v-row>
    </span>
  </span>
</template>

<script>

export default {
  name: 'Alert',
  props: {
    predefined: {
      type: Boolean,
      default: false
    },
    sets: {
      type: Array,
      default: () => [{}, {}, {}]
    },
    questions: {
      type: Array,
      default: () => [
        'What was your childhood nickname?',
        'In what city did you meet your spouse/significant other?',
        'What is the name of your favorite childhood friend?',
        'What street did you live on in third grade?',
        'What is your oldest sibling’s birthday month and year?',
        'What is the middle name of your youngest child?',
        'What is your oldest sibling\'s middle name?',
        'What school did you attend for sixth grade?',
        'What was your childhood phone number including area code?',
        'What is your oldest cousin\'s first and last name?',
        'What was the name of your first stuffed animal?',
        'In what city or town did your mother and father meet?',
        'Where were you when you had your first kiss?',
        'What is the first name of the boy or girl that you first kissed?',
        'What was the last name of your third grade teacher?',
        'In what city does your nearest sibling live?',
        'What is your youngest brother’s birthday month and year?',
        'What is your maternal grandmother\'s maiden name?',
        'In what city or town was your first job?',
        'What is the name of the place your wedding reception was held?',
        'What is the name of a college you applied to but didn\'t attend?',
        'What is the last name of your favorite high school teacher?',
        'What is your mother\'s middle name?'
      ]
    }
  },
  methods: {
    uniqueQuestion (i, sets) {
      return function (v) {
        for (let r in sets) {
          r = parseInt(r)
          if (i === r) {
            continue
          }
          if (sets[r].id === v) {
            return 'Please select a different question for each step'
          }
        }
        return true
      }
    }
  }
}
</script>
