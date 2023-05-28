import com.liferay.workspace.bundle.url.codec.BundleURLCodec;

class Main {
  private static final String BUNDLE_URL =
      "BodHRwczovL3JlbGVhc2VzLWNkbi5saWZlcmF5LmNvbS9keHAvNy40LjEzLXU3Ni9saWZlcmF5LWR4cC10b21jYXQtNy40LjEzLnU3Ni0yMDIzMDUwOTE2MDA0NDYwNy50YXIuZ3oAA=";
  private static final String RELEASE_DATE = "05/12/2023";

  public static void main(String[] args) throws Exception {
    String bundleUrlDecoded = BundleURLCodec.decode(args[0], args[1]);
    System.out.println(bundleUrlDecoded);
  }
}
