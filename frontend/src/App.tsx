import { useEffect, useMemo, useState } from 'react'
import './App.css'
import TopBar from './components/TopBar'
import type { Product, ProductsPayload, ResponsePayload, ProductByCodePayload } from './types/types'
import { BACKENDURL, ENDPOINTS, SetRequestConfig } from './utils/fetch-config'
function App() {
  const [page, setPage] = useState<number>(1)
  const [products, setProducts] = useState<Product[]>([])
  const [pageCount, setPageCount] = useState<number | null>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)

  const [searchQuery, setSearchQuery] = useState<string>('')
  const [codeResult, setCodeResult] = useState<Product | null>(null)
  const [codeLoading, setCodeLoading] = useState<boolean>(false)
  const [codeError, setCodeError] = useState<string | null>(null)

  useEffect(() => {
    const controller = new AbortController()
    const fetchProducts = async () => {
      try {
        setLoading(true)
        setError(null)
        const res = await fetch(
          `${BACKENDURL}${ENDPOINTS.productsByPage(page)}`,
          { ...SetRequestConfig('GET'), signal: controller.signal }
        )
        const json: ResponsePayload<ProductsPayload> = await res.json()
        if (!res.ok || json.error) {
          throw new Error(json.message || 'Error al obtener productos')
        }
        const data = json.data as ProductsPayload
        setProducts(data?.products ?? [])
        setPageCount(data?.page_count ?? null)
      } catch (err: any) {
        if (err.name !== 'AbortError') {
          setError(err.message ?? 'Error desconocido')
        }
      } finally {
        setLoading(false)
      }
    }
    fetchProducts()
    return () => controller.abort()
  }, [page])

  const visibleProducts = useMemo(() => {
    const q = searchQuery.trim().toLowerCase()
    if (!q || codeResult) return products
    return products.filter((p) => {
      const name = (p.product_name || '').toLowerCase()
      const brand = (p.brands || '').toLowerCase()
      const code = (p.code || '').toLowerCase()
      return name.includes(q) || brand.includes(q) || code.includes(q)
    })
  }, [products, searchQuery, codeResult])

  const handleQueryChange = (q: string) => {
    setSearchQuery(q)
    setCodeResult(null)
    setCodeError(null)
  }

  const handleCodeSearch = async (code: string) => {
    try {
      setCodeLoading(true)
      setCodeError(null)
      const res = await fetch(
        `${BACKENDURL}${ENDPOINTS.productByCode(code)}`,
        SetRequestConfig('GET')
      )
      const json: ResponsePayload<ProductByCodePayload> = await res.json()
      if (!res.ok || json.error) {
        throw new Error(json.message || 'Producto no encontrado')
      }
      const payload = json.data as ProductByCodePayload
      setCodeResult(payload.product)
    } catch (err: any) {
      setCodeError(err.message ?? 'Error en búsqueda por código')
      setCodeResult(null)
    } finally {
      setCodeLoading(false)
    }
  }

  return (
    <div className="min-h-dvh w-full bg-white text-gray-900 flex flex-col">
      <TopBar query={searchQuery} onQueryChange={handleQueryChange} onCodeSearch={handleCodeSearch} />

      <main className="flex-1 mx-auto w-full max-w-7xl px-4 sm:px-6 lg:px-8 py-8 sm:py-10">
        <section className="text-center">
          <h1 className="text-3xl sm:text-4xl font-bold tracking-tight">Bienvenido a LiquiVerde</h1>
          <p className="mt-3 text-gray-600">
            Encuentra rápidamente lo que necesitas usando la barra de búsqueda superior.
          </p>
        </section>

        <section className="mt-10 sm:mt-12">
          {loading && (
            <div className="text-center text-gray-600">Cargando productos…</div>
          )}
          {error && (
            <div className="text-center text-red-600">{error}</div>
          )}
          {!loading && !error && (
            <>
              {codeLoading && (
                <div className="text-center text-gray-600">Buscando por código…</div>
              )}
              {codeError && (
                <div className="text-center text-red-600">{codeError}</div>
              )}
              {codeResult && (
                <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
                  <article key={codeResult.code} className="rounded-lg border shadow-sm overflow-hidden bg-white">
                    <div className="aspect-[4/3] bg-gray-100">
                      {codeResult.image_url ? (
                        <img src={codeResult.image_url} alt={codeResult.product_name || codeResult.brands || 'Producto'} className="h-full w-full object-cover" loading="lazy" />
                      ) : (
                        <div className="h-full w-full flex items-center justify-center text-gray-400">Sin imagen</div>
                      )}
                    </div>
                    <div className="p-4">
                      <h3 className="font-semibold truncate" title={codeResult.product_name}>{codeResult.product_name || 'Producto sin nombre'}</h3>
                      <p className="text-sm text-gray-600 truncate" title={codeResult.brands}>{codeResult.brands || 'Marca desconocida'}</p>
                      <div className="mt-2 text-sm">Nutriscore: {codeResult.nutriscore_grade?.toUpperCase() || 'N/A'}</div>
                      <div className="mt-1 text-xs text-gray-500">Código: {codeResult.code}</div>
                    </div>
                  </article>
                </div>
              )}

              <div className="mt-6 grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
                {(!codeResult ? visibleProducts : products).map((p) => (
                  <article key={p.code} className="rounded-lg border shadow-sm overflow-hidden bg-white">
                    <div className="aspect-[4/3] bg-gray-100">
                      {p.image_url ? (
                        <img
                          src={p.image_url}
                          alt={p.product_name || p.brands || 'Producto'}
                          className="h-full w-full object-cover"
                          loading="lazy"
                        />
                      ) : (
                        <div className="h-full w-full flex items-center justify-center text-gray-400">Sin imagen</div>
                      )}
                    </div>
                    <div className="p-4">
                      <h3 className="font-semibold truncate" title={p.product_name}>{p.product_name || 'Producto sin nombre'}</h3>
                      <p className="text-sm text-gray-600 truncate" title={p.brands}>{p.brands || 'Marca desconocida'}</p>
                      <div className="mt-2 text-sm">Nutriscore: {p.nutriscore_grade?.toUpperCase() || 'N/A'}</div>
                      <div className="mt-1 text-xs text-gray-500">Código: {p.code}</div>
                    </div>
                  </article>
                ))}
                {(!codeResult ? visibleProducts.length : products.length) === 0 && (
                  <div className="col-span-full text-center text-gray-600">No hay productos para esta página.</div>
                )}
              </div>
            </>
          )}
        </section>

        <nav className="mt-12 flex items-center justify-center gap-3" aria-label="Paginación de productos">
          <button
            className="rounded-md border px-3 py-2 bg-white hover:bg-gray-50 disabled:opacity-50 hover:cursor-pointer"
            onClick={() => setPage(1)}
            disabled={page === 1}
          >
            Primero
          </button>
          <button
            className="rounded-md border px-3 py-2 bg-white hover:bg-gray-50 disabled:opacity-50 hover:cursor-pointer"
            onClick={() => setPage((p) => Math.max(1, p - 1))}
            disabled={page === 1}
          >
            Anterior
          </button>
          <span className="text-sm text-gray-700">
            Página {page}{pageCount ? ` de ${pageCount}` : ''}
          </span>
          <button
            className="rounded-md border px-3 py-2 bg-white hover:bg-gray-50 disabled:opacity-50 hover:cursor-pointer"
            onClick={() => setPage((p) => (pageCount ? Math.min(pageCount, p + 1) : p + 1))}
            disabled={pageCount !== null && page >= pageCount}
          >
            Siguiente
          </button>
          {pageCount && (
            <button
              className="rounded-md border px-3 py-2 bg-white hover:bg-gray-50 disabled:opacity-50 hover:cursor-pointer"
              onClick={() => setPage(pageCount)}
              disabled={page >= pageCount}
            >
              Último
            </button>
          )}
        </nav>
      </main>
    </div>
  )
}

export default App
