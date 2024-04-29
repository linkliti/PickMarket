import { Link } from "react-router-dom";

export default function Header() {
  const navItems = [
    { name: "Категории", path: "/categories" },
    { name: "Избранное", path: "/favorites" },
  ];

  return (
    <header className="flex w-full flex-wrap gap-5 rounded-2xl bg-sky-200 p-4">
      <Link to="/" className="flex-auto text-2xl font-extrabold">
        PickMarket
      </Link>
      <nav className="my-auto flex flex-wrap justify-between gap-5">
        {navItems.map((item) => (
          <Link key={item.name} to={item.path} className="font-bold ">
            {item.name}
          </Link>
        ))}
      </nav>
    </header>
  );
}
