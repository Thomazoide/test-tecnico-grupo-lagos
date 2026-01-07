import './App.css'
import TopBar from './components/TopBar'

function App() {
  return (
    <div className="min-h-dvh min-w-dvh bg-white text-gray-900">
      <TopBar />

      <main className="mx-auto max-w-7xl px-4 py-10">
        <section className="text-center">
          <h1 className="text-4xl font-bold tracking-tight">Bienvenido a LiquiVerde</h1>
          <p className="mt-3 text-gray-600">
            Encuentra rápidamente lo que necesitas usando la barra de búsqueda superior.
          </p>
        </section>

        <section className="mt-12 grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
          
        </section>
      </main>
    </div>
  )
}

export default App
