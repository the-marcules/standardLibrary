package test;
import org.junit.Test;

import pkg.shape.Square;

import static org.junit.Assert.assertEquals;


public class TestSquare {
  @Test
  public void testGetters() {
      double a = 1;
      Square r = new Square(a);

      assertEquals(a, r.getA(), 0);
  }
   
    @Test
   public void testArea() {
      double a = 1;
      Square s = new Square(a);

      assertEquals(1, s.area(), 0);
   }
       
       
}
