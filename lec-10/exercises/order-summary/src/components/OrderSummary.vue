<template>
  <div>
    <div class="summary-box">
      <div class="image">
        <img :src="backgroundImage" alt="hero" />
      </div>
      <div class="box-title">
        Order Summary
      </div>

      <div class="paragraph">
        You can now listen to millions of songs, audiobooks, and podcasts on any
        device anywhere you like!
      </div>
      <div class="annual-plan-box">
          <div class="music-icon">
            <img :src="musicIcon" alt="music icon" />
          </div>
          <div class="plan-change">
            <div class="annual-plan">
              {{ plan }}
            </div>
            <div class="annual-pricing">
              {{ pricing }}
            </div>
          </div>
        <button type="button" class="change" @click="changePlan"><a href="#">Change</a></button>
      </div>
      <div>
        <button class="payment" type="button" @click="showModal">Proceed To Payment</button>
        <modal v-show="isModalVisible" @close="closeModal">
          <template v-slot:body>
            Success!
          </template>
          <template v-slot:footer>
            Say yes to more shopping?
          </template>
        </modal>
        <button class="cancel" type="button">Cancel Order</button>
      </div>
    </div>
  </div>
</template>

<script>
import Modal from "./Modal";

export default {
  name: "OrderSummary",
  components: {Modal},
  data() {
    return {
      backgroundImage: require('/images/illustration-hero.svg'),
      musicIcon: require('/images/icon-music.svg'),
      isPro: false,
      isModalVisible: false,
      plan: "Annual Plan",
      pricing: "$99.99/year"
    }
  },
  methods: {
    showModal() {
      this.isModalVisible = true
    },
    closeModal() {
      this.isModalVisible = false
    },
    changePlan() {
      this.isPro = !this.isPro
      if (this.isPro) {
        this.plan = "Pro Plan"
        this.pricing = "$59.99/year"
      } else {
        this.plan = "Annual Plan"
        this.pricing = "$99.99/year"
      }
    }
  }
}
</script>

<style>

.music-icon {
  grid-area: music-icon;
  margin-left: 10px;
}

.annual-plan {
  grid-area: annual-plan;
  font-weight: 900;
  transform: translate(-50%, -50%);
  position: fixed;
  left: 50%;
  top: 72%;

}

.annual-pricing {
  grid-area: annual-pricing;
  transform: translate(-50%, -50%);
  position: fixed;
  left: 50%;
  margin-top: 28%;
}

.annual-plan, .annual-pricing {
  margin-left: -17%;
}

.change {
  grid-area: change;
  margin-right: 20px;
  border: 0;
  background-color: hsl(225, 100%, 98%);
}

.music-icon, .change {
  transform: translate(-50%, -50%);
  position: fixed;
  top: 50%;
  left: 50%;
}

.payment {
  top: 85%;
  left: 50%;
  position: fixed;
  transform: translate(-50%, -50%);
  width: 80%;
  background-color: blue;
  height: 7%;
  color: white;
  display: flex;
  margin-top: -1%;
  justify-content: center;
  align-items: center;
  border-radius: 10px;
  font-size: 97%;
  font-weight: 700;
  box-shadow: rgba(0, 0, 0, 0.35) 0 5px 15px;
}

.change:hover,
.payment:hover {
  opacity: 0.6;
}

.payment,
.cancel {
  font-size: 97%;
  font-weight: 700;

}

.cancel {
  top: 90%;
  left: 50%;
  position: fixed;
  transform: translate(-50%, -50%);
  color: gray;
  background-color: white;
  border: white;
  margin-top: 5%;
  opacity: 0.8;
}

.cancel:hover {
  opacity: 1;
}
.annual-plan-box {
  display: grid;
  /*grid-template-columns: 0.6fr 2fr 1fr;*/
  grid-template-columns: 20% 60% 10%;
  grid-template-rows: 1fr 1fr;
  gap: 0 0;
  grid-auto-flow: row;
  grid-template-areas:
    "music-icon annual-plan change"
    "music-icon annual-pricing change";
  width: 80%;
  height: 100px;
  border-radius: 10px;
  margin-top: 15%;
  left: 50%;
  position: fixed;
  transform: translate(-50%, -50%);
  background-color: hsl(225, 100%, 98%);
}
.image {
  max-width: 100%;
  height: auto;
  display: block;
  object-fit: contain;
  border-radius: 20px 20px 0 0;
  overflow:hidden;
}
.summary-box {
  position: fixed;
  margin: 23% 0 0 auto;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 450px;
  height: 680px;
  text-align: center;
  color: hsl(223, 47%, 23%);
  border-radius: 20px;
  box-shadow: rgba(0, 0, 0, 0.35) 0 5px 15px;
  background-color: white;
}

.paragraph {
  font-size: 16px;
  margin: 0 25px 15px 20px;
  text-align: center;
  line-height: 1.5em;
}

.box-title {
  font-weight: 900;
  margin: 30px 0 20px 0;
  font-size: 28px;
}

body {
  background-color: hsl(225, 100%, 94%);
}
</style>