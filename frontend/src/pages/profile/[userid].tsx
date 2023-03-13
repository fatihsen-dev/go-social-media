import { useRouter } from "next/router";

const Post = () => {
   const router = useRouter();
   const { userid } = router.query;

   return <p>Post: {userid}</p>;
};

export default Post;
