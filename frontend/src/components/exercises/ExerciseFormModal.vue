<template>
  <Transition name="modal-animation">
    <div v-show ="modalActive" class="modal">
      <Transition name="modal-animation-inner">
        <div v-show="modalActive" class="modal-inner">
          <button @click="$emit('close')">
            <CircleX/>
          </button>
          
          <slot/>
        </div>
      </Transition>
    </div>
  </Transition>

</template>

<script setup>
import { CircleX } from 'lucide-vue-next';

defineProps(['modalActive'])
defineEmits(['close']) 


</script>
