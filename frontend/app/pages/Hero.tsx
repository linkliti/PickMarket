import { Link } from "react-router-dom";

export default function Hero() {
  return (
    <>
      <section className="w-full mt-4 grow flex flex-col rounded-2xl bg-white p-4 max-md:max-w-full  items-center justify-center">
        <div className="mx-auto max-w-xl text-center">
          <h1 className="font-extrabold text-4xl">PickMarket</h1>

          <p className="mt-4 text-lg">Удобная подборка товаров</p>

          <div className="mt-8 flex flex-wrap justify-center gap-4">
            <Link to="/categories">
              <a className="block w-full rounded-lg bg-sky-200 px-5 py-3 text-md font-medium text-black sm:w-auto">
                Перейти к категориям
              </a>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
