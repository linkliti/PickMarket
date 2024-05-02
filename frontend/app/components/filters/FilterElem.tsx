import { Filter } from "@/proto/app/protos/items";
import { ReactElement } from "react";

export default function FilterElem({ filter }: { filter: Filter }): ReactElement {
  if (filter.data.oneofKind === "rangeFilter") {
    return <p>{filter.data.rangeFilter.max}</p>;
  }

  return <p>{filter.title}</p>;
}
