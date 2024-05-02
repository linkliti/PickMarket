import WhiteBlock from "@/components/base/WhiteBlock";
import { Filter } from "@/proto/app/protos/items";
import axios from "axios";
import { ReactElement, useEffect, useState } from "react";

export default function Filters({
  market,
  category,
}: {
  market: string;
  category: string;
}): ReactElement {
  const [filters, setFilters] = useState<Filter[]>([]);
  const [IsLoading, setIsLoading] = useState(true);

  async function getFilters(market: string, category: string) {
    const data = await axios.get<Filter[]>(`/api/categories/${market}/filter?url=${category}`);
    setFilters(data.data);
    return;
  }

  useEffect(() => {
    getFilters(market, category);
    setIsLoading(false);
  }, [category, market]);

  return (
    <WhiteBlock className="w-full">
      <>{IsLoading ? <div>Loading...</div> : filters.map((filter) => <div>{filter.key}</div>)}</>
    </WhiteBlock>
  );
}
