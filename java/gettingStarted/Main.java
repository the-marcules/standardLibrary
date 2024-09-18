import pkg.shape.Circle;
import pkg.shape.Rectangle;
import pkg.shape.ShapeInterface;
import pkg.shape.Square;

import java.util.ArrayList;
public class Main {
  
  public void main(String[] args) {
      ArrayList<ShapeInterface> shapes = new ArrayList<>();
      shapes.add(new Rectangle(2,2));
      shapes.add(new Circle(5));
      shapes.add(new Square(5));

      for ( ShapeInterface shape : shapes) {
        this.CalculateAll(shape);
      }
    }

    public void CalculateAll(ShapeInterface shape) {
      double area = shape.area();
      double len = shape.circumfence();

      System.out.println("Das Objekt vom Typ '" + shape.getClass() + "' hat den Umfang " + len + " und die Flaeche " + area);

    }


  }
