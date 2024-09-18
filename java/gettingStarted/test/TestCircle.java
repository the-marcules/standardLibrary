package test;

import org.junit.Test;
import static org.junit.Assert.assertEquals;

import pkg.shape.Circle;

public class TestCircle {
   
   @Test
   public void testRadius() {
      double a = 5;

      Circle c = new Circle(a);

      assertEquals(a, c.getRadius(), 0);

   }
   
   @Test
   public void testCircleArea() {
      double a = 5;

      Circle c = new Circle(a);

      assertEquals(78.5, c.area(), 0.1);
   }
   

}