import { Link } from "react-router-dom";

export default function Header() {
  const navItems = [
    { name: "Категории", path: "/categories" },
    { name: "Избранное", path: "/favorites" },
  ];

  return (
    <header className="flex w-full gap-5 whitespace-nowrap rounded-2xl bg-sky-200 px-6 py-5 font-bold text-black max-md:max-w-full max-md:flex-wrap max-md:pr-5">
      <Link to="/" className="flex-auto text-2xl">
        PickMarket
      </Link>
      <nav className="flex justify-between gap-5 self-start text-center text-base max-md:flex-wrap">
        {navItems.map((item) => (
          <Link key={item.name} to={item.path}>
            {item.name}
          </Link>
        ))}
      </nav>
    </header>
  );
}
