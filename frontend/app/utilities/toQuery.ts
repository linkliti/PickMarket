export default function filtersLink(market: number, catUrl: string): string {
  const searchParams = new URLSearchParams();
  searchParams.append("market", market.toString());
  searchParams.append("category", catUrl);
  const query = searchParams.toString();
  return `/items?${query}`;
}
