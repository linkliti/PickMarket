import BodyHeader from "@/components/base/BodyHeader";
import CategoryMarketSelector from "@/components/categories/CategoryMarketSelector";
import CategorySelect from "@/components/categories/CategorySelect";

export default function Categories() {
  return (
    <>
      <CategoryMarketSelector />
      <BodyHeader>
        <CategorySelect />
      </BodyHeader>

    </>
  );
}
