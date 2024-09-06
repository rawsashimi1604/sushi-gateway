import { HttpMethod } from "../../types/Http";
import Tag from "../../components/typography/Tag";

export interface HttpMethodTag {
  method: HttpMethod;
}

// TODO: fix colors...
function HttpMethodTag({ method }: HttpMethodTag) {
  const colorMap: { [key: string]: string } = {
    GET: "bg-green-600",
    POST: "bg-yellow-700",
    PATCH: "bg-yellow-700",
    DELETE: "bg-red-500",
    OPTIONS: "bg-blue-500",
  };

  return <Tag value={method} bgColor={colorMap[method]} />;
}

export default HttpMethodTag;
