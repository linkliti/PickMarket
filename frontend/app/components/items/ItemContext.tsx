import { Characteristic, Item } from "@/proto/app/protos/items";
import { Context, ReactElement, createContext } from "react";

export interface ItemContextType {
  item: Item;
  chars: Characteristic[];
  similar: Item[];
  totalWeight: number;
  market: number;
  maxTotalWeight: number;
  isFav: boolean;
  favUrl: string;
}
export const ItemContext: Context<ItemContextType> = createContext<ItemContextType>(
  {} as ItemContextType,
);

export function ItemContextProvider({
  children,
  data,
}: {
  children: ReactElement;
  data: ItemContextType;
}): ReactElement {
  const val = { ...data };
  return <ItemContext.Provider value={val}>{children}</ItemContext.Provider>;
}
