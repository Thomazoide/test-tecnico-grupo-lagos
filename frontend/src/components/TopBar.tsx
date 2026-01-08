import React from "react";

interface TopBarProps {
  query: string;
  onQueryChange: (q: string) => void;
  onCodeSearch: (q: string) => void;
}

function TopBar({ query, onQueryChange, onCodeSearch }: TopBarProps) {
  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const value = query.trim();
    if (value.length > 0) {
      onCodeSearch(value);
    }
  };

  return (
    <header className="w-full bg-green-600 text-white sticky top-0 z-10">
      <nav className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 py-3 flex flex-wrap items-center gap-2 sm:gap-4">
        <div className="text-lg sm:text-xl font-semibold tracking-tight">LiquiVerde</div>

        <form onSubmit={onSubmit} className="order-3 sm:order-2 w-full sm:flex-1 flex items-center gap-2 min-w-0">
          <input
            type="search"
            value={query}
            onChange={(e) => onQueryChange(e.target.value)}
            placeholder="Buscar por nombre, marca o código…"
            className="flex-1 min-w-0 rounded-md px-3 py-2 bg-green-700 text-white placeholder:text-green-200 focus:outline-none focus:ring-2 focus:ring-white/60"
            aria-label="Buscar productos"
          />
          <button type="submit" className="rounded-md bg-white text-green-700 px-3 py-2 font-medium hover:bg-green-50">Buscar</button>
        </form>

        <div className="order-2 sm:order-3 ml-auto">
          <button className="rounded-md bg-green-500/90 px-3 py-2 font-medium hover:bg-green-500">Iniciar sesión</button>
        </div>
      </nav>
    </header>
  );
}

export default TopBar;
