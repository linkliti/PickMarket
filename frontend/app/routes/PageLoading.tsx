import Loading from "@/components/base/Loading";
import WhiteBlock from "@/components/base/WhiteBlock";
import { ReactElement } from "react";

export default function PageLoading(message: string = "Загрузка..."): ReactElement {
  return (
    <WhiteBlock className="w-full">
      <Loading message={message} />
    </WhiteBlock>
  );
}
