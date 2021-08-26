<template>
  <div class="calculator">
    <table>
      <thead>
      <tr>
        <td colspan="4">
          <form id="result">
            <label for="result"></label>
            <input v-model="result" type="text" />
          </form>
        </td>
      </tr>
      </thead>
      <tbody>
      <tr>
        <td id="is-del" colspan="3"><input type="button" value="D" @click="del"/></td>
        <td class="is-clear" colspan="1"><input type="button" value="C" @click="clr"/></td>
      </tr>
      <tr>
        <td><input type="button" value="7" @click="display(7)"/> </td>
        <td><input type="button" value="8" @click="display(8)"/> </td>
        <td><input type="button" value="9" @click="display(9)"/> </td>
        <td><input type="button" value="+" @click="display('+')"/> </td>
      </tr>
      <tr>
        <td><input type="button" value="4" @click="display(4)"/> </td>
        <td><input type="button" value="5" @click="display(5)"/> </td>
        <td><input type="button" value="6" @click="display(6)"/> </td>
        <td><input type="button" value="-" @click="display('-')"> </td>
      </tr>
      <tr>
        <td><input type="button" value="1" @click="display(1)"/> </td>
        <td><input type="button" value="2" @click="display(2)"/> </td>
        <td><input type="button" value="3" @click="display(3)"/> </td>
        <td><input type="button" value="*" @click="display('*')"/> </td>
      </tr>
      <tr>
        <td class="is-zero"><input type="button" value="0" @click="display(0)"/> </td>
        <td><input type="button" value="." @click="display('.')"/> </td>
        <td class="is-equals"><input type="button" value="=" @click="equals"/> </td>
        <td><input type="button" value="/" @click="display('/')"/> </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>

export default {
  name: "Calculator",
  data() {
    return {
      result: "",
      value: {
        type: String
      }
    }
  },
  methods: {
    display(value) {
      this.result += value
    },
    del() {
      this.result = this.result.slice(0, -1)
    },
    clr() {
      this.result = ""
    },
    equals() {
      let expression = this.result
      expression = expression.replace("+", "%2b");
      console.log(expression)

      fetch('http://localhost:8000/eval?exp=' + expression, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      })
          .then(response => {
            if (!response.ok) {
              throw new Error('Network response was not ok');
            }
            return response.json();
          })
          .then(data => {
            console.log('Success:', data);
            console.log(Object.values(data)[0]);
            const returnedValue = Object.values(data)[0]
            if (!isNaN(returnedValue)) {
              this.result = returnedValue
            } else {
              this.result = "Math ERROR";
            }
          })
          .catch((error) => {
            console.error('Error: ', error)
          })
    }
  }
}
</script>

<style scoped>

</style>