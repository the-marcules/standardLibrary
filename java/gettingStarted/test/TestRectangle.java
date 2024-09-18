package test;

import org.junit.Test;
import static org.junit.Assert.assertEquals;

import pkg.shape.Rectangle;

public class TestRectangle {
   
   @Test
   public void testGetters() {
      double a = 1;
      double b = 2;

      Rectangle r = new Rectangle(a,b);

      assertEquals(a, r.getA(), 0);
      assertEquals(b, r.getB(), 0);

   }
   
   @Test
   public void testArea() {
      double a = 1;
      double b = 2;

      Rectangle r = new Rectangle(a,b);

      assertEquals(2, r.area(), 0);
   }
   

}