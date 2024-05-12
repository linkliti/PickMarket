import { PrefForm } from "@/types/filterTypes";
import { ReactElement, createContext } from "react";

export interface FilterContextType {
  categoryURL: string;
  form: PrefForm;
}
export const FilterContext = createContext<FilterContextType>({} as FilterContextType);

export function FilterContextProvider({
  children,
  data,
}: {
  children: ReactElement;
  data: FilterContextType;
}): ReactElement {
  const val = { ...data };
  return <FilterContext.Provider value={val}>{children}</FilterContext.Provider>;
}
