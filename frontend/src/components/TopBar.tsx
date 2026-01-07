import { useState } from "react";

function TopBar() {
  const [query, setQuery] = useState("");

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (query.trim().length > 0) {
      console.log("Buscar:", query);
    }
  };

  return (
    <header className="w-full bg-green-600 ">
      <nav className="mx-auto flex max-w-7xl items-center gap-4 px-4 py-3">
        <div className="text-xl font-semibold tracking-tight">LiquiVerde</div>

        <form onSubmit={onSubmit} className="flex-1 flex items-center">
          <input
            type="search"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="Buscar..."
            className="w-full max-w-xl rounded-md px-4 py-2 text-white placeholder:text-gray-200 focus:outline-none border"
            aria-label="Buscar"
          />
          <button className="ml-3" >Buscar</button>
        </form>

        <div className="ml-auto">
          <button className="text-black bg-green-400">iniciar sesión</button>
        </div>
      </nav>
    </header>
  );
}

export default TopBar;
