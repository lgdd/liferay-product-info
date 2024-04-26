import com.liferay.workspace.bundle.url.codec.BundleURLCodec;

class Main {
  public static void main(String[] args) throws Exception {
    String bundleUrlDecoded = BundleURLCodec.decode(args[0], args[1]);
    System.out.println(bundleUrlDecoded);
  }
}
