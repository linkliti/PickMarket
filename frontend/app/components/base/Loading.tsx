import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { ReactElement } from "react";

export default function Loading({ message = "Загрузка..." }: { message?: string }): ReactElement {
  return (
    <div className="flex items-center gap-2">
      <LoadingSpinner /> <p>{message}</p>
    </div>
  );
}
