<template>
    <div class="exercise-library">
        <h1>Exercise library</h1>
        <button @click="modalActive=true"> Add an exercise</button>

        <p v-if="loading">Loading...</p>

        <p v-else-if="error">{{ error }}</p>
        
        <p v-else-if="exercises.length === 0">No exercises yet</p>

        <table v-else aria-label="Exercise Library">
            <thead>
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">Name</th>
                    <th scope="col">Muscle group</th>
                    <th scope="col">Description</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(exercise, index) in exercises" :key="exercise.ID">
                    <td>{{ index + 1 }}</td>                    
                    <td>{{ exercise.Name }}</td>
                    <td>{{ exercise.MuscleGroup }}</td>
                    <td>{{ exercise.Description }}</td>
                </tr>
            </tbody>
        </table>

        <ExerciseFormModal :modalActive="modalActive" @close="modalActive=false">
            <div class="modal-content">
                <h1> Add an exercise</h1>
                <form @submit.prevent="addExercise">
                    <label>
                        Exercise Name
                    <input v-model="newExercise.name" type="text" placeholder="e.g. Bench Press" />
                    </label>
                    <label>
                        Muscle Group
                    <input v-model="newExercise.muscle_group" type="text" placeholder="e.g. Chest" />
                    </label>
                    <label>
                        Description
                    <input v-model="newExercise.description" type="text"/>
                    </label>
                    <button type="submit">Save</button>
                </form>
            </div>
        </ExerciseFormModal>

    </div>

</template>
<script setup>
import{ref, onMounted} from 'vue'
import ExerciseFormModal from '@/components/exercises/ExerciseFormModal.vue'
import {createExercise, listExercises} from '@/api/exercises.js'

const newExercise = ref({ name: '', muscle_group: '', description: ''})
const modalActive = ref(false)
const exercises = ref([])
const loading = ref(false)
const error = ref(null)

onMounted ( async() =>{
    loading.value = true
    try{
        exercises.value = await listExercises()
    } catch (err){
        error.value = err.message
    } finally{
        loading.value=false
    }
})

async function addExercise() {
    try{
        await createExercise(newExercise.value)
        exercises.value = await listExercises()
        modalActive.value = false
        newExercise.value = { name: '', muscle_group: '', description: '' }
    }catch(err){
        error.value = err.message
    }
}

</script>
<style lang="scss" scoped>
.exercise-library{
    background: var(--paper);
}


</style>