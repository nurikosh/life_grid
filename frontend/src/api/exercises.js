const BASE = import.meta.env.VITE_API_URL || 'http://localhost:8080'

async function handleResponse(response) {
  if (!response.ok) {
    const message = await response.text()
    throw new Error(message)
  }
  return response.json()
}

export async function listExercises() {
    const res = await fetch(`${BASE}/exercises`)
    return handleResponse(res)
}

export async function createExercise(payload) {
  const res = await fetch(`${BASE}/exercises`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })
  return handleResponse(res)
}

export async function getExerciseById(id) {
  const res = await fetch(`${BASE}/exercises/${id}`)
  return handleResponse(res)
}

export async function deleteExercise(id) {
  const res = await fetch(`${BASE}/exercises/${id}`, { method: 'DELETE' })
  if (!res.ok) throw new Error(await res.text())
}