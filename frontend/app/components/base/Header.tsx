import { Link } from "react-router-dom";

function Header() {
  const navItems = [
    { name: "Категории", path: "/categories" },
    { name: "Избранное", path: "/favorites" },
  ];

  return (
    <header className="flex gap-5 px-6 py-5 w-full font-bold text-black whitespace-nowrap bg-sky-200 rounded-2xl max-md:flex-wrap max-md:pr-5 max-md:max-w-full">
      <Link to="/" className="flex-auto text-2xl">
        PickMarket
      </Link>
      <nav className="flex gap-5 justify-between self-start text-base text-center">
        {navItems.map((item) => (
          <Link key={item.name} to={item.path}>
            {item.name}
          </Link>
        ))}
      </nav>
    </header>
  );
}

export default Header;
