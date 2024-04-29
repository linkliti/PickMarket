import BodyHeader from "@/components/base/BodyHeader";
import CategoryMarketSelector from "@/components/categories/CategoryMarketSelector";
import CategorySelect from "@/components/categories/CategorySelect";
import { ReactElement } from "react";

export default function Categories(): ReactElement {
  return (
    <>
      <CategoryMarketSelector />
      <BodyHeader>
        <CategorySelect />
      </BodyHeader>

    </>
  );
}
